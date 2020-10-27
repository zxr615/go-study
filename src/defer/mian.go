package main

import "fmt"

func printError() {
	fmt.Println("兜底执行")
}

func divide()  {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var i = 1
	var j = 0

	if j == 0 {
		panic("除数不能为 0 ！")
	}

	var k = i / j

	fmt.Printf("%d / %d = %d\n", i, j, k)
}

func main() {
	divide()
	fmt.Println("main body")
}
