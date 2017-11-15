package main

import "fmt"

func main() {
	for i := 0; i < 500; i++ {
		// fmt.Println(i, "-", string(i), "-", []byte(string(i)))
		// []byte("hello")
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
