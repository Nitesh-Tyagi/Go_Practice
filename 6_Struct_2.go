package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {

	// GOLANG DOES NOT AHVE INHERITANCE
	// INSTEAD HAS COMPOSITION

	// INHERITANCE -> IS A MODEL
	// COMPOSITION -> HAS LIKE CHARACTERISTICS ( EMBEDDING )

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Print(b, "\n")

}
