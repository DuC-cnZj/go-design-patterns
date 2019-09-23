package strategy

import "fmt"

type sayer interface {
	say()
}

type Jack struct {
}

func (jack Jack) say() {
	fmt.Println("jack saying cool.")
}

type Jone struct {
}

func (jone Jone) say() {
	fmt.Println("jone saying bad.")
}
