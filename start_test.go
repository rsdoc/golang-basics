package basicsgo

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	variadicFn()
	fmt.Println("------------------------")
	unfurlingSliceFn()
	fmt.Println("------------------------")
	deferGo()
	fmt.Println("------------------------")
	recieverFn()
	fmt.Println("------------------------")
	interfacesInGo()
}
