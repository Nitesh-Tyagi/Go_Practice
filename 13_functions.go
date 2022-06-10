package main

import (
	"fmt"
)

func helloOnce(name *string) {
	fmt.Print("Hello ", *name, "\n")
	*name = *name + "!"
}

func helloMultiple(name *string, times int) {
	fmt.Print("\n")
	for times > 0 {
		// fmt.Print(times, ".  Hello ", name, "\n")
		helloOnce(name)
		times--
	}
	fmt.Print("\n")
}

func main() {
	var name string = "Nitesh Tyagi"

	helloOnce(&name)

	helloMultiple(&name, 5)
}
