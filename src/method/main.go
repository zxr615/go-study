package main

import (
	"fmt"
	"math"
)

func main() {
	r1 := Rectangle{4, 5}
	r2 := Circle{2}

	fmt.Println("r1 rectangle", r1.area())
	fmt.Println("r1 circle", r2.area())
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.width
}

func (r Circle) area() float64 {
	return r.radius * r.radius * math.Pi
}

//func area(r Rectangle) float64 {
//	return r.width * r.height
//}

