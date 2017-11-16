package main

import "fmt"

func main() {
	average := average(3, 5, 7, 8, 4, 12, 43, 11, 36, 28)
	fmt.Println(average)
}

func average(number ...float64) float64 {
	fmt.Println(number)        // slice of float64
	fmt.Printf("%T\n", number) // Type of slice
	var total float64
	for _, v := range number {
		total += v
	}
	return total / float64(len(number))
}
