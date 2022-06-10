package main

import (
	"fmt"
)

func main() {
	/*
		fmt.Print("Start\n")
		defer fmt.Print("Middle\n")
		fmt.Print("End\n")
	*/

	/*
		// DEFER WORKS IN LAST IN FIRST OUT
		// DEFER IS USED MOSTLY TO CLOSE OUT RESOURCES, SO LIFO IS USEFUL
		defer fmt.Print("Start\n")
		defer fmt.Print("Middle\n")
		defer fmt.Print("End\n")
	*/

	// DEFER TAKES USES VALUES GIVEN WHEN THE DEFER STATEMENT IS CALLED NOT WHEN EXECUTED
	a := "start"
	defer fmt.Print(a, "\n")
	a = "end"
}
