package main

import "fmt"

type person struct {
	name        string
	age, weight int
	height      float32
}

func main() {
	p1 := person{"John", 24, 50, 1.5}
	fmt.Printf("p1: %v\n", p1)

	p2 := person{
		name:   "Rob",
		age:    29,
		height: 1.8,
		weight: 50,
	}
	fmt.Printf("p2: %v\n", p2)

	p3 := person{}
	fmt.Println(p3)
	p3.age = 30
	p3.name = "Lol"
	p3.weight = 70
	p3.height = 1.9
	fmt.Printf("p3: %v\n", p3)
}
