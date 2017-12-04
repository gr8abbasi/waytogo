package main

import (
	"fmt"
)

func main() {
	gen := gen(1000)
	// out := make(chan int)
	for c := range factorial(gen) {
		fmt.Println(c)
	}
}

func gen(n int) <-chan int64 {
	out := make(chan int64)
	go func() {
		for index := 0; index < n; index++ {
			out <- int64(index)
		}
		close(out)
	}()
	return out
}

func factorial(n <-chan int64) <-chan int64 {
	out := make(chan int64)
	for c := range n {
		go func(num int64) {
			var total int64
			total = 1
			for index := num; index > 0; index-- {
				total *= index
			}
			out <- total
		}(c)
	}
	return out
}

// func merge(cs <-chan int) <-chan int {
// 	out := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(len(n))
// 	for c := range n {
// 		go func(num int) {
// 			total := 1
// 			for index := num; index > 0; index-- {
// 				total *= index
// 			}
// 			out <- total
// 			wg.Done()
// 		}(c)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()

// 	return out
// }
