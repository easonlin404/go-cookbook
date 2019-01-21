package main

import "fmt"

type S []int

func (s S) Get() string {
	if s == nil {
		return "it's nil"
	} else {
		return "hello"
	}

}

func main() {
	var ss []int

	if ss == nil {
		fmt.Println("ss nil")
	}

	fmt.Println(len(ss))
	s := getSlice()

	if s == nil {
		fmt.Println("nil")
		fmt.Printf("123:%s\n", s.Get())
	} else {
		fmt.Println("qqq")
	}

	var m *my
	if m == nil {
		fmt.Println("m == nil")
	}
	fmt.Println(m.Get())

}

func getSlice() S {
	return nil
}

type my struct {
}

func (s *my) Get() string {
	if s == nil {
		return "it's nil"
	} else {
		return "mymymymymymym"
	}

}
