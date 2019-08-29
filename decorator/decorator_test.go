package decorator

import (
	"fmt"
)

func ExampleDecor() {
	f := func() {
		fmt.Println("base")
	}

	Decor(Decor(f))()

	// Output:
	// a
	// a
	// base
	// b
	// b
}

func ExampleDecor2() {
	base := BaseColor{"red"}
	d := ColorDecor{baseColor: base}
	fmt.Println(base.GetColor())
	fmt.Println(d.GetColor())
	// Output:
	// red
	// red  decorator
}
