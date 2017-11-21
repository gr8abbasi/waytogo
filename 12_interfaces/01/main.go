package main

import (
	"fmt"
	"math"
)

// Square Implements Shape Interface
type Square struct {
	side float64
}

// Circle Implements Shape Interface
type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Shape Interface
type Shape interface {
	area() float64
}

func info(shape Shape) {
	fmt.Println(shape)
	fmt.Println(shape.area())
}

func main() {
	square := Square{5}
	circle := Circle{3}

	info(square)
	info(circle)
}
