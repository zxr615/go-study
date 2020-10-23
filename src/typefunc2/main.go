package main

import "fmt"

func main() {
	var a Integer = 2

	if a.Equal(3) {
		fmt.Println(a, " = 2")
	} else {
		fmt.Println("!=")
	}
}

type Integer int

func (a Integer) Equal (b Integer) bool {
	return a == b
}


