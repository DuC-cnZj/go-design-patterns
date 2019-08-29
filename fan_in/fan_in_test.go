package fan_in

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	c1 := make(chan int, 3)
	c2 := make(chan int, 3)

	c1 <- 1
	c1 <- 2
	c1 <- 3

	c2 <- 10
	c2 <- 11
	c2 <- 12

	close(c1)
	close(c2)

	out := Merge(c1, c2)

	for value := range out {
		fmt.Println(value)
	}
}
