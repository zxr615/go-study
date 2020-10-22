package main

import "fmt"

func main() {
	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	var a, b []byte

	a = ar[5:6]
	//a[0] = 'z'
	b = ar[2:5]

	c := cap(a)
	l := len(a)
	fmt.Printf("a = %s \n", a)
	fmt.Printf("b = %s \n", b)
	fmt.Printf("a cap = %v \n", c)
	fmt.Printf("l = %v \n", l)
}