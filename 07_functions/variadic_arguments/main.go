package main

import "fmt"

func main() {
	data := []float64{2, 3, 9, 8, 34, 56, 73, 23, 45}
	average := average(data...)
	fmt.Println(average)
}
func average(number ...float64) float64 {
	var total float64
	for _, v := range number {
		total += v
	}
	return total / (float64(len(number)))
}
