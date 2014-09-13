package main

import "fmt"

func bark() {
	fmt.Println("Woof!")
}

func retTwoInts() (int, int) {
	return 1, 2
}

func addTwoInts(a, b int) int {
	return a + b
}

func adder(nos ...int) (sum int) {
	for _, n := range nos {
		sum += n
	}
	return
}

func main() {
	bark()

	fmt.Printf("4 + 5 = %d\n", addTwoInts(4, 5))

	a, b := retTwoInts()
	fmt.Printf("a: %d, b: %d\n", a, b)

	fmt.Printf("sum of first five nos: %d\n", adder(1, 2, 3, 4, 5))
}
