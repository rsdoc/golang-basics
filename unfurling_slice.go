package basicsgo

import "fmt"

func unfurlingSliceFn() {
	fmt.Println("Testing unfurlingSliceFn()")

	xi := []int{10, 20, 30}

	// unfurling xi and passing to variadic fns
	x := multiply(xi...)

	fmt.Println("The value after multiplication is : ", x)

	fmt.Println("Exit from unfurlingSliceFn()")
}


// accepting int slice as an argument
func multiply(x ...int) int {
	mul := 1

	for _, v := range x {
		mul *= v
	}

	return mul
}