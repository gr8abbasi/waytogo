package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string)
	// myMap := map[string]string{}
	// var myMap map[string]string // This will return nil map and cannot use append method on maps
	myMap["K1"] = "Kashif"
	myMap["K2"] = "Abbasi"
	fmt.Println(myMap)
}
