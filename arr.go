package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr[2] = 1
	arr[4] = 3
	fmt.Println(arr)

	arr2 := [5]int{0, 0, 1, 0, 3}
	fmt.Println(arr2)
}
