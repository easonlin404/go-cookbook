package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{
		Name: "Eason",
		Age:  10,
	}

	o := []interface{}{"123", 3, p}
	b, err := json.Marshal(o)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Println(string(b))
}
