package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)              // Number of waitgroup e.g wait for 2 func to execute
	go incrementor("Foo:") // with go it doesn't display anything we have to use WaitGroup
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter // demostration we can do counter++
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond) // random time
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.g
