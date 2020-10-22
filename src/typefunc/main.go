package main

import "fmt"

func main() {
	slice := []int {1, 2, 3, 4, 5, 7}
	add := filter(slice, isOdd)

	fmt.Println(add)
}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer % 2 == 0{
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var rs []int
	for _, v := range slice {
		if f(v) {
			rs = append(rs, v)
		}
	}

	return rs
}

