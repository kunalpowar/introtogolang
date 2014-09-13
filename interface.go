package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Color() string
}

func printTotalArea(shapes ...Shape) {
	var area float32
	for _, shape := range shapes {
		area += shape.Area()
	}
	fmt.Printf("Total area of shapes is %v\n", area)
}

type square struct {
	side  float32
	color string
}

func (s *square) Area() float32 {
	return s.side * s.side
}

func (s *square) Color() string {
	return s.color
}

type circle struct {
	r     float32
	color string
}

func (c *circle) Area() float32 {
	return math.Pi * c.r * c.r
}

func (c *circle) Color() string {
	return c.color
}

func printAnything(a interface{}) {
	fmt.Println(a)
}

func (s *square) String() string {
	return fmt.Sprintf("sq: side: %v, color: %v", s.side, s.color)
}

func main() {
	a := square{1, "red"}
	b := circle{2.5, "yellow"}

	fmt.Println(&a)

	printTotalArea(&a, &b)

	printAnything("hello")
	c := []int{1, 2, 3}
	printAnything(c)
}
