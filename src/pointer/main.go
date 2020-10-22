package main

import "fmt"

func main() {
	 i := 3

	 fmt.Println("i = ", i)

	 i = add(&i);
	 fmt.Println("i+1 = ", i)
	 fmt.Println("i = ", i)
}

func add(a *int) int {
	*a = *a +1
	return *a
}
