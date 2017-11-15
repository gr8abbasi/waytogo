package main

import "fmt"

// Rune is one character, int32 (4 bytes), represented by ''
// String Literal is collection of Rune
// Raw String Literal is represented by ``
// Interpreted String is represented by ""

func main() {
	for i := 0; i < 500; i++ {
		// fmt.Println(i, "-", string(i), "-", []byte(string(i)))
		// []byte("hello")
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))

		// Rune
		a := 'a'
		fmt.Println(a) // 97

		b := "b"
		fmt.Println(b) // b
	}
}
