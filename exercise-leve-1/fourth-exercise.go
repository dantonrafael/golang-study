package main

import (
	"fmt"
)

type danton int

var x danton

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // %T is like type() function in python
	x = 42
	fmt.Println(x)
}
