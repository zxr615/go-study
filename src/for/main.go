package main

import (
	"fmt"
)

func main() {

	numbers := make(map[string]int)
	numbers["age"] = 18
	numbers["stature"] = 170
	numbers["vision"] = 5

	for k, v := range numbers {
		fmt.Println(k,"=>",v)
	}

	for _, v := range numbers{
		fmt.Println("v", v)
	}

	sum := 0;
	for i:=0; i < 5; i++ {
		sum++
		fmt.Println("sum = ", sum)
	}

}
