package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()

	fmt.Println("cupnum = ", cpuNum)
}
