package main

import (
	"fmt"
)

func main() {
	var x [58]int // array always have fixed size otherwise slice
	fmt.Println(x)
	fmt.Println(len(x))
	x[42] = 999
	fmt.Println(x[42])
}
