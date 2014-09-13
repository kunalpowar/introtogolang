package main

import (
	"fmt"
)

type person struct {
	name        string
	age, weight int
	height      float32
}

func (p *person) PrintDetails() {
	fmt.Printf("person -> name: %v, age: %d, weight: %d, height: %f\n", p.name, p.age, p.weight, p.height)
}

func main() {
	p1 := person{"John", 24, 50, 1.5}
	p1.PrintDetails()
}
