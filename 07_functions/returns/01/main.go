package main

import "fmt"

func main() {
	fmt.Println(greet("Foo", "Bar")) // Arguments
}

func greet(fname, lname string) string { // Params/Parameters
	return fmt.Sprint(fname, lname)
}
