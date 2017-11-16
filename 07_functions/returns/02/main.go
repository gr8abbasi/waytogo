package main

import "fmt"

func main() {
	fmt.Println(greet("Foo", "Bar"))
}

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
