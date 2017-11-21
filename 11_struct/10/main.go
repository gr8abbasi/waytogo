package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Person Decode JSON
type Person struct {
	FirstName  string
	LastName   string
	Age        int
	gender     string // unexported fields doesn't work with Decode
	Profession string `json:"Beruf"` // Tag to change field name while converting to JSON
	Phone      int    `json:"-"`     // Tag to remove field from JSON
}

func main() {
	var p1 Person

	fmt.Println(p1)

	reader := strings.NewReader(`{"FirstName":"Kashif","LastName":"Abbasi","Age":32, "gender":"Male","Beruf":"Software Engineer", "Phone":"1234567890"}`)

	json.NewDecoder(reader).Decode(&p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.FirstName, "=", p1.LastName, "=", p1.Age, "=", p1.gender, "=", p1.Profession, "=", p1.Phone)
}
