package main

import "fmt"

func main() {
	person := Person{}
	person.Exercise("篮球")
}

type Body interface {
	Speak(content string)
	Exercise(item string)
}

type Person struct {
}

func (person Person) Speak(content string) {
	fmt.Println("speak: ", content)
}

func (person Person) Exercise(item string) {
	fmt.Println(" Exercise: ", item)
}

