package a

import (
	"fmt"

	"x04/b"
)

type A struct {
}

func (a A) PrintA() {
	fmt.Println(a)
}

func NewA() *A {
	a := new(A)
	return a
}

func RequireB() {
	o := b.NewB()
	o.PrintB()
}
