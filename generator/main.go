package main

import "fmt"

// 使用了无缓存管道
func Count(start int, end int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i <= end ; i++ {
			// Blocks on the operation
			// 阻塞
			ch <- i
		}

		close(ch)
	}(ch)

	return ch
}

func main() {
	for i := range Count(1, 100) {
		fmt.Println(i)
	}
}
