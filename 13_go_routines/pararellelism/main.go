package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// init() is required prior to 1.5 its just to explain
// this is used for pararellelism
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
}

func main() {
	wg.Add(2) // Number of waitgroup e.g wait for 2 func to execute
	go foo()  // with go it doesn't display anything we have to use WaitGroup
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond) // delay to see concurrency and mixed printout
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}
