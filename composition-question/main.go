package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

//unmark this part
// func (t *Teacher) ShowA() {
// 	fmt.Println("teacher showA")
// }

func main() {
	t := Teacher{}
	t.ShowA()
}
