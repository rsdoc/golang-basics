package basicsgo

import "fmt"

func variadicFn() {
	fmt.Println("Testing  variadicFn()")

	x := sum(1, 2, 3, 4, 5, 6)

	fmt.Println("Sum of all is : ", x)

	fmt.Println("Exit from variadicFn()")
}

// we can pass any number of arguments to variadic functions
// lexical elements docs - operators and delimiters
func sum(x ...int) int {

	sum := 0

	// here x will be the int slice
	for _, v := range x {
		sum += v
	}

	return sum
}
