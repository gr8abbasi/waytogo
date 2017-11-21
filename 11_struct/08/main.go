package main

import (
	"encoding/json"
	"fmt"
)

// Person UnMarshal JSON
type Person struct {
	FirstName  string
	LastName   string
	Age        int
	gender     string // unexported fields doesn't work with UnMarshal
	Profession string `json:"Beruf"` // Tag to change field name while converting to JSON
	Phone      int    `json:"-"`     // Tag to remove field from JSON
}

func main() {
	var p1 Person

	fmt.Println(p1)

	bs := []byte(`{"FirstName":"Kashif","LastName":"Abbasi","Age":32,"Beruf":"Software Engineer", "Phone":"1234567890"}`)

	json.Unmarshal(bs, &p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.FirstName, "=", p1.LastName, "=", p1.Age, "=", p1.Profession, "=", p1.Phone)
}
