package main

import "fmt"

func main() {
	a := 40
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a // var b *int = &a
	// Referencing
	fmt.Println(b)
	// Dereferencing
	fmt.Println(*b)

	*b = 39
	fmt.Println(a)
}
