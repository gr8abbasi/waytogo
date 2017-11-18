package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 1, 43)
	fmt.Println(slice)
	changeMe(slice)
	fmt.Println(slice)
}

// Slice, Map and Channel are already Reference Type no need to use &
func changeMe(x []string) {
	x[0] = "Kashif"
}
