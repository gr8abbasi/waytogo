package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough // run
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough // run
	case "Julian":
		fmt.Println("Wassup Julian") // run
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
}
