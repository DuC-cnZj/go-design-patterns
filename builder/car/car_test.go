package car

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type RedCar struct {
	color  Color
	wheels Wheels
	speed  Speed
}

func (r *RedCar) Wheels(wheels Wheels) Builder {
	r.wheels = wheels

	return r
}

func (r *RedCar) TopSpeed(speed Speed) Builder {
	r.speed = speed

	return r
}

func (r *RedCar) Build() Interface {
	return r
}

func (r *RedCar) Color(color Color) Builder {
	r.color = color

	return r
}

func (r *RedCar) Drive() error {
	fmt.Printf("%s的车%s轮子最高速度%f跑起来了", r.color, r.wheels, r.speed)
	return nil
}
func (r *RedCar) Stop() error {
	return nil
}

func TestBuilder(t *testing.T) {
	redCar := RedCar{}

	r := redCar.
		Wheels(SportsWheels).
		Color(RedColor).
		TopSpeed(KPH).
		Build()

	res := r.Drive()

	assert.Nil(t, res)
}
