package main

import (
	"fmt"
)

func main() {
	var mySlice = make([]int, 0, 5)
	fmt.Println("---------------------------")
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("---------------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println(len(mySlice), "-", cap(mySlice), "-", mySlice[i])
	}
}
