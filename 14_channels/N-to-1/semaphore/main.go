package main

import (
	"fmt"
)

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	// N-to-1
	// will lauch 10/n go routines
	for index := 0; index < n; index++ {
		go func() {
			for index := 0; index < 10; index++ {
				c <- index
			}
			done <- true
		}()
	}

	// offloading n times
	go func() {
		for index := 0; index < 10; index++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
