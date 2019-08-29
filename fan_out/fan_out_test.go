package fan_out

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSplit(t *testing.T) {
	c := make(chan int, 10)
	wg := sync.WaitGroup{}

	// 分成几个管道
	num := 10

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
			time.Sleep(time.Millisecond)
		}
		close(c)
	}()

	cs := Split(c, num)

	wg.Add(num)

	for i := 0; i < num; i++ {
		fmt.Println("启动一个协程", i)
		go func(i int) {
			for range cs[i] {
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}
