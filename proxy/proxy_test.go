package proxy

func ExampleCurl() {
	ab := Ab{}

	proxy := AbProxy{ab: ab}

	proxy.curl("/duc")
	proxy.curl("/ab")

	// Output:
	// ab abProxy curl url= /duc
	// ab curl url= /duc
	// ab abProxy curl url= /ab
}
