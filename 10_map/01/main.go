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
	fmt.Println(myMap) // map[2:James 3:Bond 0:Hello 1:Mr.]

	myMap[4] = "007"   // Add key
	fmt.Println(myMap) // map[3:Bond 4:007 0:Hello 1:Mr. 2:James]

	myMap[1] = "Mr"    // Update key
	fmt.Println(myMap) // map[3:Bond 4:007 0:Hello 1:Mr 2:James]

	delete(myMap, 2)   // Delete key
	fmt.Println(myMap) // map[3:Bond 4:007 0:Hello 1:Mr]
}
