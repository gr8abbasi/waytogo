package main

import (
	"fmt"
)

func main() {
	// Conversion: used for types except interface
	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	// Assertion: only used for interface
	var val interface{} = 7
	fmt.Printf("%T\n", val)
	fmt.Printf("%T\n", val.(int)) // int(val) will throw error
}
