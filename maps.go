package main

import "fmt"

func main() {
	myMap := make(map[string]int)
	myMap["a"] = 1
	myMap["12345"] = 2

	fmt.Println(myMap)

	delete(myMap, "12345")
	fmt.Println(myMap)

	a := myMap["a"]
	fmt.Printf("myMap[\"a\"]: %v\n", a)

	_, ok := myMap["a"]
	fmt.Printf("myMap[\"a\"] exists: %v\n", ok)
}
