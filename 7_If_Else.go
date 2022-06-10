package main

import (
	"fmt"
)

func main() {
	a := 55
	if a < 10 {
		fmt.Print("Small\n")
	} else if a > 10 {
		fmt.Print("Large\n")
	} else {
		fmt.Print("Equal\n")
	}
}
