package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p := New(30)

	times := 0
	for {
		m := <-p
		num := m.GetNum()
		fmt.Println("pool num is=", num)
		p <- m

		time.Sleep(time.Millisecond)

		if num == 1 {
			times++
		}

		if times == 3{
			break
		}
	}
}
