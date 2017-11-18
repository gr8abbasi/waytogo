package main

import (
	"fmt"
)

func main() {
	records := make([][]string, 0)

	student1 := make([]string, 3)
	student1[0] = "Foo Student"
	student1[1] = "90%"
	student1[2] = "80%"
	records = append(records, student1)

	student2 := make([]string, 3)
	student2[0] = "Bar Student"
	student2[1] = "70%"
	student2[2] = "85%"
	records = append(records, student2)

	fmt.Println(records)
}
