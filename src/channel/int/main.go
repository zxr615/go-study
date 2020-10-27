package main

import "fmt"

func main() {
	// 定义 channel
	var intChan chan int
	intChan = make(chan int, 10)

	a := 1

	// 向 channel 写入数据
	intChan<- a

	// 从 channel 取出数据
	b := <-intChan

	fmt.Println("b = ", b)
}
