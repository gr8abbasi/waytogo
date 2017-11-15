package main

import "fmt"

func main() {
	x := 0
	for {
		x++
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)
		if x >= 50 {
			break
		}
	}
}
