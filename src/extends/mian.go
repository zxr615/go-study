package main

type Animal struct {
	name string
}

func (animal Animal) Call() string {
	return "叫声"
}

func (animal Animal) Food() string {
	return "食物"
}

func (animal *Animal) SetName(name string) {
	animal.name = name
}

func (animal Animal) GetName() string {
	return animal.name
}


type Dog struct {
	Animal
}

func (dog Dog) Call() string {
	return "旺旺旺"
}

func (dog Dog) Food() string {
	return "骨头"
}

func main() {
	dog := Dog{}
	dog.SetName("Snoopy")

	println("name = ", dog.GetName())
	println("call = ", dog.Call())
	println("food = ", dog.Food())
}





