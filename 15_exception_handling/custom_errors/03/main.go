package main

import (
	"fmt"
	"log"
)

// MySqrtError Struct
type MySqrtError struct {
	number float64
	err    error
}

func (mse *MySqrtError) Error() string {
	return fmt.Sprintf("Error: something went wrong here: %v", mse.err)
}

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt function
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("Negative Number: Square root of negative number: %v", f)
		return 0, &MySqrtError{f, err}
	}
	// Some Implementation here
	return 42, nil
}
