package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	byt := []byte(`{"name":"eason"}`)
	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("data: %v\n", data)
}
