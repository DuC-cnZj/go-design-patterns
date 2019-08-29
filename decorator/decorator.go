package decorator

import "fmt"

type Decorator func()

func Decor(f Decorator) Decorator {
	return func() {
		fmt.Println("a")
		f()
		fmt.Println("b")
	}
}

type Color interface {
	GetColor() string
}

type BaseColor struct {
	color string
}

func (c BaseColor) GetColor() string {
	return c.color
}

type ColorDecor struct {
	baseColor Color
}

func (c ColorDecor) GetColor() string {
	return c.baseColor.GetColor() + "  decorator"
}
