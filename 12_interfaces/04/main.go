package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

// With Value Type Reciever
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// With Pointer Type Reciever
// func (c *Circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

func info(shape Shape) {
	fmt.Println(shape.area())
}

// Receivers       Values
// -----------------------------------------------
// (t T)           T and *T
// (t *T) 		  *T

func main() {
	c := Circle{7}
	info(c) // or info(&c) // Value Type Reciever

	c1 := Circle{5}
	info(&c1) // Pointer Type Reciever
}
