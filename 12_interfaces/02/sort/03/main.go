package main

import (
	"fmt"
	"sort"
)

// Implements Interface interface of Sort Package
type people struct {
	Name string
	Age  int
}

func (p people) String() string {
	return fmt.Sprintf("YAYAYA %s: %d", p.Name, p.Age)
}

// ByAge Type
type ByAge []people

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	// Slice of String of Type People
	people := []people{
		{"Zeno", 31},
		{"John", 35},
		{"Jenny", 23},
	}

	fmt.Println(people[0])
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
