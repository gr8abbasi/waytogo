package main

import (
	"fmt"
)

func main() {
	age := make(map[string]int)
	fmt.Println(age)
	changeMe(age)
	fmt.Println(age)
}

// Slice, Map and Channel are already Reference Type no need to use &
func changeMe(m map[string]int) {
	m["Kashif"] = 32
}
