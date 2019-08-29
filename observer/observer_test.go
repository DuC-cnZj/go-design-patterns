package observer

func ExampleObserver() {
	s := MySubject{events: make([]Observer, 0)}
	a := AObserver{}
	b := BObserver{}

	s.attach(a)
	s.attach(b)

	s.notify()

	// Output:
	// a observer do sth ....
	// b observer do sth ....
}
