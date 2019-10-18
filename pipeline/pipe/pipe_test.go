package pipe

import (
	"fmt"
)

func ExamplePipeline() {
	p := &Pipeline{
		middlewares: nil,
	}

	p.through(func(m myFunc) myFunc {
		return func(i ...interface{}) {
			fmt.Println("in")
			m(i...)
			fmt.Println("out")
		}
	}).send(func(i ...interface{}) {
		fmt.Printf("核心逻辑：%s，收到的参数：%v\n", "core", i)
	})(1,2,3)

	// Output:
	// in
	// 核心逻辑：core，收到的参数：[1 2 3]
	// out
}