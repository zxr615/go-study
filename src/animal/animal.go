package animal

type Animal struct {
	name string
}

func NewAnimal(name string) *Animal {
	return &Animal{name:name}
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