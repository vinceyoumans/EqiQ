package b

import (
	"fmt"

	"x05/i"
)

type B struct {
}

func (b B) PrintB() {
	fmt.Println(b)
}

func NewB() *B {
	b := new(B)
	return b
}

func RequireA(o i.Aprinter) {
	fmt.Println("--------   from B ----------")
	o.PrintA()
}
