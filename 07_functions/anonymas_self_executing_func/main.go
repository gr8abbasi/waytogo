package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("I'm Anonymas Self Executing Funct!")
	}() // parantheses used for self executing
}
