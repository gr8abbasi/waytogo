package main

import "fmt"

type contact struct {
	greetings string
	name      string
}

func changeByType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	changeByType(7)
	changeByType("Whose there?")
	var t = contact{"good morning", "Tom"}
	changeByType(t)
	changeByType(t.greetings)
	changeByType(t.name)
}
