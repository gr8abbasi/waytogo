package main

import (
	"errors"
	"log"
)

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt function
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Negative Number: Square root of negative number")
	}
	// Some Implementation here
	return 42, nil
}
