package main

import (
	"fmt"
)

func main() {
	defer world() // defers execuation of func untill the func exit in which it is
	hello()
	// Hello World!
}

func world() {
	fmt.Print("World!\n")
}

func hello() {
	fmt.Print("Hello ")
}
