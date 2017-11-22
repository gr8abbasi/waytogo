package main

import (
	"fmt"
)

type animal struct {
	Sound string
}

type cat struct {
	animal
	friendly bool
}

type dog struct {
	animal
	annoying bool
}

// Using Empty Interface for func() Param
func characteristics(ch interface{}) {
	fmt.Println(ch)
}

func main() {
	cat := cat{animal{"Mewo"}, true}
	dog := dog{animal{"Woof"}, true}

	characteristics(cat)
	characteristics(dog)
}
