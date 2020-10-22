package main

import "fmt"

func main() {

	var a [10]int
	var b [11]int
	fmt.Printf("a type = %T \n", a)
	fmt.Printf("b type = %T \n", b)


	i := [2]int{1,2}
	j := [2]int{3,4}
	//f := [3]int{5, 6, 7}

	k := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	z := [2][2]int{i, j}
	//p := [2][3]int{i, f}


	fmt.Printf("i type = %v \n", i)
	fmt.Printf("j type = %v \n", j)
	fmt.Printf("k type = %v \n", k)
	fmt.Printf("z type = %v \n", z)
	//fmt.Printf("p type = %v \n", p)


}
