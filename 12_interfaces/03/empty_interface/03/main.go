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

func main() {
	cat := cat{animal{"Mewo"}, true}
	dog := dog{animal{"Woof"}, true}

	critter := []interface{}{cat, dog}
	fmt.Println(critter)
}
