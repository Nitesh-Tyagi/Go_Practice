package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	/*
		var a int = 10
		var b *int = &a
		fmt.Print("Value of a : ", a, "\n")
		fmt.Print("Address of a / Valaue of b : ", &a, " , ", b, "\n")
		fmt.Print("Value at address stored in b / Value at *b : ", *b, "\n") // * IS DEREFERENCE OPERATOR
		fmt.Print("\n")

		// CHANGING VALUE USING DEREFERENCE OPERATOR
		*b = 15
		fmt.Print(a, " _ ", *b, "\n")
		fmt.Print("\n")

		// GO DOES NOT ALLOW SIMPLE POINTER ARITHMETIC UNLIKE C,C++
	*/

	// /*
	// UNDECLARED POINTERS HOLD NULL VALUE
	var ms *myStruct
	fmt.Print(ms, "\n")
	ms = new(myStruct)
	fmt.Print(ms, "\n")
	(*ms).foo = 42
	fmt.Print(ms.foo, "\n")
	ms.foo = 35 // THIS IS COMPILER HELPING US OUT, WRITTEN IN A WAY TO ACCEPT THIS ALSO, TO MAKE IT EASIER FOR US
	fmt.Print(*ms, "\n")

	// */

	// SLICES AND MAPS ARE BUILT ON POINTERS

}
