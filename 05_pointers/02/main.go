package main

import "fmt"

func changeX(x *int) {
	fmt.Println(x)
	*x = 0
}

func main() {
	x := 5

	fmt.Println(x)
	fmt.Println(&x)

	// Go is always pass by value, here passing memory address value
	changeX(&x)

	fmt.Println(x)
}
