package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetOutput(createFile())
	openFile()
}

func createFile() (file *os.File) {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func openFile() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatal(err)
	}
}
