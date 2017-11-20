package main

import (
	"fmt"
)

// Person Pointers on Struct
type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := &Person{
		Name: "Kashif Abbasi",
		Age:  32,
	}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.Name, "===", p1.Age)
}
