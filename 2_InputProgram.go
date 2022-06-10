package main

import (
	"fmt"
)

func main() {
	var fname string
	fmt.Print("Enter your firstname : ")
	fmt.Scan(&fname)

	var lname string
	fmt.Print("Enter your lastname : ")
	fmt.Scan(&lname)

	fmt.Print("Your full name is : " + fname + " " + lname)
}
