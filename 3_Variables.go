package main

import (
	"fmt"
)

var (
	actorName    string = "Elizabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	// WAYS TO DECLARE VARIABLE - 1
	// var i int
	// i = 42

	// WAYS TO DECLARE VARIABLE - 2
	// var i int = 42

	// WAYS TO DECLARE VARIABLE - 3
	// i := 42

	i := 25
	j := float32(i)
	fmt.Print(i)
	fmt.Printf("\n%v, %T", i, i)
	fmt.Printf("\n%v, %T\n", j, j)

	// BINARY OPERATORS
	// 	& AND
	//	| OR
	// 	^ XOR
	// 	&^ AND NOT (NOT)
	// 	(a<<3) BIT SHIFT LEFT 3 PLACES
	//  (a>>3) BIT SHIFT RIGHT 3 PLACES

	a := 8
	fmt.Println(a << 2)

	z := 3 + 2i
	fmt.Printf("\n%v : %T\n", z, z)
	fmt.Printf("%v : %T\n", real(z), real(z))
	fmt.Printf("%v : %T\n", imag(z), imag(z))

}
