package main

import (
	"fmt"
)

func main() {
	greet := greeting()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}

func greeting() func() string {
	return func() string {
		return "Hello Func Expression from Closure!"
	}
}
