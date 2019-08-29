package factory

type MyFactory int

const (
	AppleFactory MyFactory = 1 << iota
	BananaFactory
	PeachFactory
)

type Fruit interface {
	Get() string
}

type apple struct {
}

func (apple *apple) Get() string {
	return "apple"
}

type peach struct {
}

func (apple *peach) Get() string {
	return "peach"
}

type banana struct {
}

func (apple *banana) Get() string {
	return "banana"
}

func Create(factoryType MyFactory) Fruit {
	switch factoryType {
	case AppleFactory:
		return new(apple)
	case PeachFactory:
		return new(peach)
	case BananaFactory:
		return new(banana)
	default:
		panic("error type")
	}
}
