package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

/*
  Will caused as following error message:

  runtime: goroutine stack exceeds 1000000000-byte limit
 fatal error: stack overflow
*/
func main() {
	p := &People{}
	p.String()
}
