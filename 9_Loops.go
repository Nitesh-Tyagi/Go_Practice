package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")

	for i, j := 0, 5; i < 5; i, j = i+1, j-1 {
		fmt.Print(i, "_", j, "\n")
	}
	fmt.Print("\n")

	a := 0
	// FOR CAN ALSO BE USED AS WHILE LOOP
	for a < 5 {
		fmt.Print(a, "_")
		a++
	}
	fmt.Print("\n\n")

	// INFINITE OR UNDETERMINED LOOP
	for {
		if a == 0 {
			break
		}
		fmt.Print(a, "->")
		a--
	}
	fmt.Print("\n\n")

A:
	for {
		if a == 7 {
			break A
			// USE LABEL TO BREAK OUT OF MULTIPLE LOOPS
			// DEFINE LABEL BEFORE HAND, BREAK TO LABEL
		}
		fmt.Print(a, "->")
		a++
	}
	fmt.Print("\n\n")

	// FOR RANGE LOOP OVER COLLECTIONS
	// SLICE, ARRAY, MAP, STRING, CHANNELS
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Print(k, " : ", v, "\n")
	}
	fmt.Print("\n")

	// FOR RANGE LOOP - ONLY VALUES
	for _, v := range s {
		fmt.Print(v, "\n")
	}
	fmt.Print("\n")

	// FOR RANGE LOOP - ONLY KEY
	for k := range s {
		fmt.Print(k, " : ", s[k], "\n")
	}
	fmt.Print("\n")
}
