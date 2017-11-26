package main

import (
	"fmt"
)

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for index := 0; index < 100000; index++ {
			c <- index
		}
		close(c)
	}()

	for index := 0; index < n; index++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	for index := 0; index < n; index++ {
		<-done
	}
}
