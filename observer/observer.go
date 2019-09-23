package observer

import "fmt"

type Subject interface {
	attach(o Observer)
	detach(o Observer)
	notify()
}

type Observer interface {
	update()
}

type MySubject struct {
	events []Observer
}

func (s *MySubject) attach(o Observer) {
	s.events = append(s.events, o)
}

func (s *MySubject) detach(o Observer) {

}

func (s *MySubject) notify() {
	for _, e := range s.events {
		e.update()
	}
}

type AObserver struct {
}

func (o AObserver) update() {
	fmt.Println("a observer do sth ....")
}

type BObserver struct {
}

func (o BObserver) update() {
	fmt.Println("b observer do sth ....")

}
