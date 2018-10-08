# go-cookbook

* [Convert GO struct to JSON](#convert-go-struct-to-json)
* [Convert JSON to GO struct](#convert-json-to-go-struct)
* [Pass slice as a function argument](#pass-slice-as-a-function-argument)



# Question
* [Composition in Go](https://github.com/easonlin404/go-cookbook/blob/master/composition-question/main.go)
* [Interface in Go](https://github.com/easonlin404/go-cookbook/blob/master/interface-question/main.go)

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

## Convert JSON to GO struct
See the [example code](https://github.com/easonlin404/go-cookbook/blob/master/convert-json-to-go-struct/main.go).

```go
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

```

## Pass slice as a function argument
See the [example code](https://github.com/easonlin404/go-cookbook/blob/master/pass-slice-as-a-function-argument/main.go).
```go
package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	appendSlice(s)
	fmt.Println(s)

	modifySlice(s)
	fmt.Println(s)
}

func appendSlice(s []int) {
	s = append(s, 4)
}

func modifySlice(s []int) {
	s[0] = 0
}

```