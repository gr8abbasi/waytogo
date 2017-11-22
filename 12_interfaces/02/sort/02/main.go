package main

import (
	"fmt"
	"sort"
)

func main() {
	// Reverse Slice of String
	students := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.Reverse(sort.StringSlice(students)))
	fmt.Println("-------------------------------------")
	fmt.Println(students)

	// Reverse Slice of Integer
	integers := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(integers)))
	fmt.Println("-------------------------------------")
	fmt.Println(integers)

}
