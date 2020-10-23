package main

import "fmt"

func main() {
	student := CreateStudent()
	student.SetName("fanwei")
	student.SetAge(10)

	fmt.Println("GetName: ", student.GetName())
	fmt.Println("GetName: ", student.GetAge())
}

type Student struct {
	name string
	age uint8
}

func NewStudent(name string, age uint8) *Student {
	rs := &Student{name, age}
	return rs
}

func CreateStudent() *Student {
	return &Student{}
}

func (student *Student) SetName(name string)  {
	student.name = name
}

func (student Student) GetName() string {
	return student.name
}

func (student *Student) SetAge(age uint8)  {
	student.age = age
}

func (student Student) GetAge() uint8 {
	return student.age
}
