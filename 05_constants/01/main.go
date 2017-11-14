package main

import "fmt"

// Uses bitwise operator and move 10 bits and so on
const (
	_  = iota             // 0 (which will be thrown away i.e. _ blank identifier)
	KB = 1 << (iota * 10) // (1 * 10)
	MB = 1 << (iota * 10) // (2 * 10)
	GB = 1 << (iota * 10) // (3 * 10)
	TB = 1 << (iota * 10) // (4 * 10)
)

func main() {
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
