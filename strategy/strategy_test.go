package strategy

func ExampleStrategy() {
	jack := Jack{}
	jack.say()

	jone := Jone{}
	jone.say()

	// Output:
	// jack saying cool.
	// jone saying bad.
}
