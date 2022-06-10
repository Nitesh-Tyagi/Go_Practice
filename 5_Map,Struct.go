package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	name       string
	companions []string
}

func main() {
	pop := make(map[string]int, 10)
	pop = map[string]int{
		"abc": 55,
		"def": 38,
		"ghi": 45,
	}
	fmt.Print(pop["a"])
	fmt.Print("\n\n")

	aDoctor := Doctor{
		number:     3,
		name:       "Jon Pertwee",
		companions: []string{"Liz Shaw", "Jo Grant"},
	}

	// ANONYMUS STRUCT
	bActor := struct{ name string }{name: "Bill Hader"}

	fmt.Print(bActor, "\n")

	fmt.Print(aDoctor, "\n")
}
