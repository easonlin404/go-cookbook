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
