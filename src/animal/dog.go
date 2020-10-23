package animal

type Dog struct {
	*Animal
}

func (dog Dog) Call() string {
	return "旺旺旺"
}

func (dog Dog) Food() string {
	return "骨头"
}
