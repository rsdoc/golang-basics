package basicsgo

import "fmt"

type person struct {
	fn string
	ln string
}

type secretAgent struct {
	person
	ltk bool
}

func recieverFn() {
	sa1 := secretAgent{
		person: person{
			"Rahul",
			"Singh",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
}

func (sa secretAgent) speak() {
	fmt.Println("Hello from secret agent", sa.fn+" "+sa.ln)
}
