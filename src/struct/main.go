package main

import "fmt"

func main() {

	var person Person
	person.name = "fanwei"
	person.age = 18
	person.phone = "13800138000"

	//fanwei := Student{Person{"fanwei", 18, "13800138000"}, "Computer"}
	fanwei := Student{person, "computer", "13800138001"}
	fmt.Println(fanwei.Person.phone)
	fmt.Println(fanwei.phone)

}

type Person struct {
	name string
	age uint8
	phone string
}

type Student struct {
	Person
	speciality string
	phone string
}


