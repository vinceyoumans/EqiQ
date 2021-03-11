package a

import (
	"fmt"

	"x05/b"
)

type A struct {
}

func (a A) PrintA() {
	fmt.Println("--------   from A ----------")
	fmt.Println(a)
}

func NewA() *A {
	fmt.Println("-------- Creating new A ----------")
	a := new(A)
	return a
}

func RequireB() {
	fmt.Println("-------- in A RequireB ----------")
	o := b.NewB()
	o.PrintB()
}
