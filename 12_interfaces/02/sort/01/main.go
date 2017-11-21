package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	// Slice of String of Type People
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)

	// Slice of String
	students := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(students)
	fmt.Println("-------------------------------------")
	fmt.Println(students)

	// Slice of Integer
	integers := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(integers)
	fmt.Println("-------------------------------------")
	fmt.Println(integers)

}
