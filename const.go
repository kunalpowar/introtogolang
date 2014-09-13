package main

import "fmt"

const (
	stringConst = "stringconst"
	intConst    = 2
)

func main() {
	var aStr string
	aStr = stringConst
	fmt.Printf("aStr: %s\n", aStr)

	defTypeStr := "a default string"
	fmt.Printf("defTypeStr: %s, type is %T\n", defTypeStr, defTypeStr)

	defTypeInt := 1
	fmt.Printf("defTypeInt: %d, type is %T\n", defTypeInt, defTypeInt)

	fmt.Printf("%v type is %T\n", "Hello", "Hello")
	fmt.Printf("%v type is %T\n", stringConst, stringConst)
	fmt.Printf("%v type is %T\n", 0, 0)
	fmt.Printf("%f type is %T\n", 0.0, 0.0)
	fmt.Printf("%c type is %T\n", 'x', 'x')
	fmt.Printf("%c type is %T\n", '0', '0')
	fmt.Printf("%v type is %T\n", 0i, 0i)
}
