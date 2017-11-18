package main

import (
	"fmt"
)

func main() {
	myMap := map[int]string{
		0: "Hello",
		1: "Mr.",
		2: "James",
		3: "Bond",
	}

	//delete(myMap, 2)

	if val, exists := myMap[2]; exists { // called as "Comma OK Idiom"
		fmt.Println("Value is: ", val)
		fmt.Println("Exists is: ", exists)
	} else {
		fmt.Println("That value doesn't exists!")
	}
}
