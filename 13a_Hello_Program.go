package main

import (
	"fmt"
)

func upgrade(a string) string {
	if a[0] >= 'a' && a[0] <= 'z' {
		return string('A' + (int(a[0]) - int('a')))
	} else {
		return string(a[0])
	}
}

func display(name *string, pos int) {
	(*name) = (*name)[:pos] + upgrade(string((*name)[pos])) + (*name)[pos+1:]
	fmt.Print(*name, "\n")
}

func hello(name string) {
	for k := range name {
		fmt.Print(k+1, ".  Hello ")
		display(&name, k)
	}
}
func main() {
	name := "Nitesh Tyagi"
	hello(name)
}
