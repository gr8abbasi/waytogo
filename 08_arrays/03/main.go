package main

import (
	"fmt"
)

func main() {
	var x [256]int
	fmt.Println(x)
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = i
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if v > 50 {
			break
		}
	}
}
