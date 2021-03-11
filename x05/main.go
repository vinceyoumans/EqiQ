package main

import (
	"fmt"
	"x05/a"
)

func main() {
	fmt.Println("--------   from main ----------")
	o := a.NewA()
	o.PrintA()
}
