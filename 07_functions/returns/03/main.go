package main

import "fmt"

func main() {
	fmt.Println(greet("Foo", "Bar"))
}

func greet(fname string, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
