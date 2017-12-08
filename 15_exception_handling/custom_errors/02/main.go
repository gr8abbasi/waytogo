package main

import (
	"fmt"
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
		//MySqrtError := errors.New("Negative Number: Square root of negative number")
		MySqrtError := fmt.Errorf("Negative Number: Square root of negative number: %v", f)
		return 0, MySqrtError
	}
	// Some Implementation here
	return 42, nil
}
