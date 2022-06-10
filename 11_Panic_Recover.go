package main

import (
	"fmt"
)

func main() {
	/*
		// PANIC IS BASICALLY ERROR
		// BCZ PROGRAM CANNOT RUN ANY FURTHER AND PANICING
		a, b := 1, 0
		ans := a / b
		fmt.Print(ans, "\n")
	*/

	/*
		fmt.Print("Start\n")
		panic("Something bad happened\n")
		fmt.Print("End\n")
	*/

	/*
		// DEFER STILL GETS EXECUTED BEFORE PANIC HAPPENS
		// ASSURES ALL RESOURCES ARE CLOSED BEFORE PROGRAM IS KILLED
		fmt.Print("Start\n")
		defer fmt.Print("This was deferred\n")
		panic("Something bad happened\n")
		fmt.Print("End\n")
	*/

	// /*
	// RECOVER MAY BE USED TO TRY TO RECOVER PROGRAM WHEN PROGRAM IS PANICKING
	// HERE AN ANONYMUS FUNCTION DEFERS TO END AND IF PANIC OCCURS,
	// IT LETS US KNOW DURING RECOVERY
	fmt.Print("Start\n")
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("Program is Panicking!\n")
		} else {
			fmt.Print("Program is fine\n")
		}
	}()
	panic("Something bad happened\n")
	fmt.Print("End\n")
	// */
}
