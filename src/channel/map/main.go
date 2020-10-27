package main

import "fmt"

func main() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 3)


	m1 := make(map[string]string, 20)
	m1["name"] = "fanwei"
	m1["age"] = "1"


	mapChan<- m1

	m1s := <-mapChan

	fmt.Println(m1s["name"])

}
