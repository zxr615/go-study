package main

import "fmt"

func main() {

	a := Max(1, 1);
	b, c := SumAndSubtract(1, 1)
	d, e := SumAndSubtract2(1,1)

	Arg(1,3,5,8)

	println("max = " , a)
	println("SS= " , b, c)
	println("SS2= " , d, e)

}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func SumAndSubtract(a, b int) (int, int) {
	sum := a + b
	subtract := a - b

	return sum, subtract
}

func SumAndSubtract2(a, b int) (sum int, subtract int) {
	sum = a + b
	subtract = a - b

	return
}

func Arg(arg ...int)  {
	fmt.Println("arg[0] =", arg[0])
	for k, v := range arg{
		fmt.Println("arg", k, "=>", v)
	}
}