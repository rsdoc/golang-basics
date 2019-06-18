package basicsgo

import "fmt"

func deferGo() {
	defer coo()
	doo()
}

func coo() {
	fmt.Println("I am deferred and in coo()")
}

func doo() {
	fmt.Println("I am in doo()")
}
