package main

// Person Promotions/Overriding of Methods
type Person struct {
	Name string
	Age  int
}

// DoubleZero Promotions/Overriding of Methods
type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) greetings() {
	println("I'm a regular Person Greetings!")
}

func (d DoubleZero) greetings() {
	println("I'm DoubleZero Greetings!")
}

func main() {
	p1 := Person{
		Name: "Kashif Abbasi",
		Age:  32,
	}

	p2 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age:  45,
		},
		LicenseToKill: false,
	}

	p1.greetings()        // specifically called Person method
	p2.greetings()        // Overriden Person method
	p2.Person.greetings() // specifically called Person method
}
