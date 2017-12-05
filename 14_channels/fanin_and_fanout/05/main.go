package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for s := range incrementor(2) {
		fmt.Println(s)
	}
}

func incrementor(n int) <-chan string {
	out := make(chan string)
	done := make(chan bool)
	for index := 0; index < n; index++ {
		go func(n int) {
			for i := 0; i < 200; i++ {
				time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond) // random time
				out <- fmt.Sprintf("%d %s %d", n, " Print Process Number:", i)
			}
			done <- true
		}(index)
	}

	go func() {
		for index := 0; index < 2; index++ {
			<-done
		}
		close(out)
	}()

	return out
}
