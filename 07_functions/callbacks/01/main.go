package main

import "fmt"

func main() {
	visitor([]int{2, 3, 4, 5, 6}, func(n int) {
		fmt.Println(n)
	})
}

func visitor(n []int, callback func(n int)) {
	for _, v := range n {
		callback(v)
	}
}
