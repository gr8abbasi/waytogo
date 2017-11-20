package main

import (
	"fmt"
)

// Person Embbeded
type Person struct {
	firstName string
	lastName  string
	Gender    string
	Age       int
}

// DoubleZero Embbeding Types
type DoubleZero struct {
	Person
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			firstName: "Kashif",
			lastName:  "Abbasi",
			Gender:    "Male",
			Age:       32,
		},
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			firstName: "Jame",
			lastName:  "Bond",
			Gender:    "Male",
			Age:       45,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1.lastName, "-", p1.LicenseToKill)
	fmt.Println(p2.lastName, "-", p2.LicenseToKill)
}
