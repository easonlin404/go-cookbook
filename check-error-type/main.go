package main

import "fmt"

var actulMyError = &myError{code: 500, msg: "actulMyError"}

func main() {
	if err := oops(); err != nil {
		if e, ok := err.(*myError); ok {
			fmt.Println(e.code)
			fmt.Println(e.msg)
		}
	}

	if err := oops(); err != nil {
		//_ = err.(*myError2)

		//panic: interface conversion: error is *main.myError, not *main.myError2
	}

	if err := oops2(); err == actulMyError {
		fmt.Println("err == actulMyError")
	} else {
		fmt.Println("xxxx")
	}

}

func oops() error {
	return &myError{code: 1, msg: "ooops"}
}

func oops2() error {
	return actulMyError
}

type myError struct {
	code int
	msg  string
}

func (e *myError) Error() string {
	return fmt.Sprintf("code :%d msg:%s", e.code, e.msg)
}

type myError2 struct {
	code int
	msg  string
}

func (e *myError2) Error() string {
	return fmt.Sprintf("code :%d msg:%s", e.code, e.msg)
}
