package main

import (
	"fmt"
)

type person struct { //Encapsulation i.e. hiding other type by one type e.g. person
	firstName string
	lastName  string
	Gender    string
	Age       int
}

func (p person) fullName() string {
	return p.firstName + p.lastName
}

func main() {
	p1 := person{"Kashif", "Abbasi", "Male", 32}
	p2 := person{"Jame", "Bond", "Male", 45}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
