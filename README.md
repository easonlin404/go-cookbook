# go-cookbook

* [Convert GO struct to JSON](#Convert-GO-struct-to-JSON)



## Convert GO struct to JSON
See the [example code](https://github.com/easonlin404/go-cookbook/blob/master/convert-go-struct-to-json/main.go).
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