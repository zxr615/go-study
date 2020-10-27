package main

import (
	"fmt"
	"sync"
	"time"
)

var rsMap = make(map[int]uint64)
var lock sync.Mutex

func main() {
	num := 20

	//normal(num)
	gof(num)
	time.Sleep(time.Second * 1)

	for k, v := range rsMap {
		fmt.Println(k, "的阶乘是", v)
	}
}

func gof(num int)  {
	for i := 1; i <= num; i++ {
		go fact(i)
	}
}


func normal(num int)  {
	for i := 1; i <= num; i++ {
		fact(i)
	}
}

func fact(n int) {
	// 3! = 1*2*3
	var rs uint64 = 1
	for i := 1; i <=n ; i++ {
		rs = rs * uint64(i)
	}
	lock.Lock()
	rsMap[n] = rs
	lock.Unlock()
}