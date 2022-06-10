package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

const (
	_ = iota
	a2
	b2
)

func main() {
	fmt.Print(a, b, c)
	fmt.Print("\n")
	fmt.Print(a2, b2)
}
