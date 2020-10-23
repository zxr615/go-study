package main

import (
	"animal"
	"fmt"
)

func main() {
	ani := animal.NewAnimal("Snoopy")
	dog := animal.Dog{Animal: ani}

	fmt.Println(dog.GetName(), dog.Call(), dog.Food())
}