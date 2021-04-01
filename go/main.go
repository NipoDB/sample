package main

import (
	"github.com/NipoDB/nipolib"
	"strconv"
	"os"
	"runtime"
	"sync"
	"fmt"
)
var Wait sync.WaitGroup
var Lock sync.Mutex

func set(config *nipolib.Config, regex string, max int) {
	adas:=0
	for n:=0 ; n <= max ; n++ {
		adas ++
		if adas == 1000 {
			fmt.Println(n)
			adas = 0
		}
		key := regex + "_" + strconv.Itoa(n)
		nipolib.Set(config, key, strconv.Itoa(n))
	}
}
func get(config *nipolib.Config, regex string, max int) {
	adas:=0
	for n:=0 ; n <= max ; n++ {
		adas ++
		if adas == 1000 {
			fmt.Println(n)
			adas = 0
		}
		key := regex + "_" + strconv.Itoa(n)
		nipolib.Get(config, key)
	}
}
func main() {
	// token server port get/set regex count cores threads
	config := nipolib.CreateConfig(os.Args[1], os.Args[2], os.Args[3])
	regex := os.Args[5]
	max,_ := strconv.Atoi(os.Args[6])
	cores,_ := strconv.Atoi(os.Args[7])
	threads,_ := strconv.Atoi(os.Args[8])

	if os.Args[4] == "set" {
		runtime.GOMAXPROCS(cores)
		for thread := 0; thread < threads; thread++ {
			Wait.Add(1)
			go func() {
    	    	defer Wait.Done()
					Lock.Lock()
					set(config, regex, max)
					Lock.Unlock()
			}()
		}
		Wait.Wait()
		result,_ := nipolib.Avg(config, regex+"_.*")
		fmt.Println(result)
	}
	if os.Args[4] == "get" {
		runtime.GOMAXPROCS(cores)
		for thread := 0; thread < threads; thread++ {
			Wait.Add(1)
			go func() {
    	    	defer Wait.Done()
					Lock.Lock()
					get(config, regex, max)
					Lock.Unlock()
			}()
		}
		Wait.Wait()
		result,_ := nipolib.Avg(config, regex+"_.*")
		fmt.Println(result)
	}
}
