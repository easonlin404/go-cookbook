# go-cookbook

* [Convert GO struct to JSON](#Convert-GO-struct-to-JSON)



## Convert GO struct to JSON
See the [example code]().
```go
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
	p := &Person{
		Name: "Eason",
		Age:  10,
	}
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

```