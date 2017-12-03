package main

import (
	"fmt"
	"time"
)

func main() {
	c := fanIn(boring("Tom"), boring("Jerry"))
	for index := 0; index < 10; index++ {
		fmt.Println(<-c)
	}
	fmt.Println("You both are boring! I'm leaving!")
}

func boring(msg string) <-chan string {
	out := make(chan string)
	go func() {
		for index := 0; ; index++ {
			out <- fmt.Sprintf("%s %d", msg, index)
			time.Sleep(5)
		}
	}()
	return out
}

func fanIn(input1, input2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			out <- <-input1
		}
	}()
	go func() {
		for {
			out <- <-input2
		}
	}()
	return out
}
