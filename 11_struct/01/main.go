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

func main() {
	p1 := person{"Kashif", "Abbasi", "Male", 32}
	p2 := person{"Jame", "Bond", "Male", 45}
	fmt.Println(p1.firstName, "-", p1.lastName, "-", p1.Gender, "-", p1.Age)
	fmt.Println(p2.firstName, "-", p2.lastName, "-", p2.Gender, "-", p2.Age)
}
