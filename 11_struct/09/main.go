package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person Encode JSON
type Person struct {
	FirstName  string
	LastName   string
	Age        int
	gender     string
	Profession string `json:"Beruf"` // Tag to change field name while converting to JSON
	Phone      int    `json:"-"`     // Tag to remove field from JSON
}

func main() {
	p1 := Person{
		FirstName:  "Kashif",
		LastName:   "Abbasi",
		Age:        32,
		gender:     "Male",
		Profession: "Software Engineer",
		Phone:      1234567890,
	}
	fmt.Println(p1)

	json.NewEncoder(os.Stdout).Encode(p1)
}
