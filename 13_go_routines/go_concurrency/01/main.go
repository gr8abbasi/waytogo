package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

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
