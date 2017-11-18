package main

import (
	"fmt"
)

func main() {
	age := 44
	fmt.Println(age) // 44
	changeMe(&age)
	fmt.Println(age) // 32
}

func changeMe(x *int) {
	*x = 32
	fmt.Println(&x)
}
