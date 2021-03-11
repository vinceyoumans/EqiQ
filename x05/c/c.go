package c

import (
	"fmt"
	"x05/a"
	"x05/b"
)

func PrintC() {
	fmt.Println("--------   from C ----------")
	o := a.NewA()
	b.RequireA(o)
}
