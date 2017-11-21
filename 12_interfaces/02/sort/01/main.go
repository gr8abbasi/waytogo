package main

import (
	"fmt"
	"sort"
)

// Implements Interface interface of Sort Package
type people []string

func (a people) Len() int           { return len(a) }
func (a people) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a people) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	// Slice of String of Type People
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	// sort.Strings(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	// Slice of String
	students := []string{"Zeno", "John", "Al", "Jenny"}
	//sort.Sort(sort.StringSlice(students))
	sort.StringSlice(students).Sort()
	//sort.Strings(students)
	fmt.Println("-------------------------------------")
	fmt.Println(students)

	// Slice of Integer
	integers := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	//sort.Ints(integers)
	//sort.IntSlice.Sort(integers)
	sort.Sort(sort.IntSlice(integers))
	fmt.Println("-------------------------------------")
	fmt.Println(integers)

}
