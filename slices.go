package main

import "fmt"

func main() {
	arr := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("arr: %v type: %T\n", arr, arr)

	slc := arr[0:2]
	fmt.Printf("slc: %v type: %T\n", slc, slc)

	slc2 := []int{1, 2, 3, 4}
	fmt.Printf("slc2: %v\n", slc2)

	slc3 := make([]int, 5)
	fmt.Printf("slc3: %v\n", slc3)

	slc4 := make([]int, 5, 10)
	fmt.Printf("slc4: %v\n", slc4)

	slc5 := append(slc4, 11)
	fmt.Printf("slc5: %v\n", slc5)

	slc6 := make([]int, 2)
	copy(slc6, arr[:])
	fmt.Printf("slc6: %v\n", slc6)
}
