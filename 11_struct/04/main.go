package main

import (
	"fmt"
)

// Person Promotions/Overriding of Fields
type Person struct {
	firstName string
	lastName  string
	Gender    string
	Age       int
}

// DoubleZero Promotions/Overriding of Fields
type DoubleZero struct {
	Person
	firstName     string
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
		firstName:     "I will override Kashif!",
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			firstName: "James",
			lastName:  "Bond",
			Gender:    "Male",
			Age:       45,
		},
		firstName:     "I will override James!",
		LicenseToKill: false,
	}

	fmt.Println(p1.firstName, "-", p1.Person.firstName)
	fmt.Println(p2.firstName, "-", p2.Person.firstName)
}
