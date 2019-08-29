// see https://books.studygolang.com/go-patterns/messaging/fan_out.html
package fan_out

import "fmt"

// Split a channel into n channels that receive messages in a round-robin fashion.
func Split(ch <-chan int, n int) []chan int {
	fmt.Printf("把一个 chan 分成 %d 个chan\n", n)

	cs := make([]chan int, 0)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	// Distributes the work in a round robin fashion among the stated number
	// of channels until the main channel has been closed. In that case, close
	// all channels and return.
	distributeToChannels := func(ch <-chan int, cs []chan int) {
		// Close every channel when the execution ends.
		defer func(cs []chan int) {
			for _, c := range cs {
				close(c)
			}
		}(cs)

		for {
			for index, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						fmt.Println("break")
						return
					}

					fmt.Printf("管道 %d 接收到了值 %d\n", index, val)

					c <- val
				}
			}
		}
	}

	go distributeToChannels(ch, cs)

	return cs
}
