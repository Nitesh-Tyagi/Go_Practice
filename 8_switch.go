package main

import (
	"fmt"
)

func main() {
	// TAG SYNTAX - PASSING VALUE b TO SWITCH STATEMENT
	switch b := 2 + 3; b {
	case 1:
		fmt.Print("one\n")
	case 2, 4, 8:
		fmt.Print("power of two\n")
	default:
		fmt.Print("not one or power of two\n")
	}
	fmt.Print("\n")

	a := 10
	// TAGLESS SYNTAX - NOT PASSING VALUE a TO SWITCH STATEMENT
	switch {
	case a <= 10:
		fmt.Print("less than or equal to ten\n")
		fallthrough
		// BREAK STATEMENT IS IMPLIED
		// BREAK MAY BE USED TO SHORT CIRCUIT THE STATEMENT IN CASE OF ERROR MAYBE
		// FALLTHROUGH STATEMENT IS NOT IMPLIED
	case a <= 20:
		fmt.Print("less than or equal to twenty\n")
	default:
		fmt.Print("greater than twenty\n")
	}
	fmt.Print("\n")

	// USE OF INTERFACE
	// INTERFACE : DATA TYPE THAT ACCEPTS ALL DATA TYPE VALUES (INT, FLOAT, STRING, ARRAY ...)
	var c interface{} = 3.45
	switch c.(type) {
	case int:
		fmt.Print("Is an Integer\n")
	case float64:
		fmt.Print("Is a float64\n")
	case string:
		fmt.Print("Is a string\n")
	default:
		fmt.Print("Is another datatype\n")
	}
	fmt.Print("\n")
}
