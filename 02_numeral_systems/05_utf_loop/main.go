package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#x \t %o \t %q", i, i, i, i, i)
	}
}
