package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)              // Number of waitgroup e.g wait for 2 func to execute
	go incrementor("Foo:") // with go it doesn't display anything we have to use WaitGroup
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond) // random time
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter)) // access without race
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.g
