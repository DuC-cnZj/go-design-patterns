package proxy

import "fmt"

type Curler interface {
	curl(url string)
}

type Ab struct {
	url string
}

func (ab Ab) curl(url string) {
	fmt.Println("ab curl url=", url)
}

type AbProxy struct {
	ab Curler
}

func (abProxy AbProxy) curl(url string) {
	fmt.Println("ab abProxy curl url=", url)

	if url == "/duc" {
		abProxy.ab.curl(url)
	}
}
