// see https://books.studygolang.com/go-patterns/messaging/fan_in.html
package fan_in

import (
	"sync"
)

// Merge different channels in one channel
func Merge(cs ...<-chan int) <-chan int{
	var wg sync.WaitGroup

	// 阻塞
	out := make(chan int)

	// Start an send goroutine for each input channel in cs. send
	// copies values from c to out until c is closed, then calls wg.Done.
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	// Start a goroutine to close out once all the send goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
