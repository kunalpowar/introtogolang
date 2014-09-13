package main

import (
	"math/rand"
	"time"

	"fmt"
)

func f(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s : %d\n", name, i)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}
func main() {
	go f("foo")
	go f("bar")
	var i string
	fmt.Scanln(&i)
}
