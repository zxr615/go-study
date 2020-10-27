package main

import (
	"fmt"
	"math/rand"
)

func main() {
	personChan := make(chan Person, 11)

	for i := 0; i < 10; i++ {
		//person := Person{Name: "FANWEI", Age: 1, Address: "ch"}
		person := Person{}
		person.Name = "Cat " + string(i)
		person.Age = rand.Intn(20)
		person.Address = "猫星球"

		personChan<- person
	}

	//fmt.Println(<-personChan)

	for i := 0; i < 10; i++ {
		rs := <-personChan
		fmt.Println(rs)
	}
}


type Person struct {
	Name string
	Age int
	Address string
}
