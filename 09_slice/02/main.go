package main

import (
	"fmt"
)

func main() {
	sliceOne := []string{"Monday", "Tuseday"}
	sliceTwo := []string{"Wednesday", "Thursday", "Friday"}

	sliceOne = append(sliceOne, sliceTwo...) // Append item
	fmt.Println(sliceOne)
	sliceOne = append(sliceOne[:2], sliceOne[3:]...) //Delete Item
	fmt.Println(sliceOne)
}
