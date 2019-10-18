package main

import (
	"fmt"
	"net/http"
	"time"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

type Pipeline struct {
	middlewares []middleware
}

func (p *Pipeline) through(m ...middleware) *Pipeline {
	p.middlewares = append(p.middlewares, m...)

	return p
}

func (p *Pipeline) send(handler http.HandlerFunc) http.HandlerFunc {
	length := len(p.middlewares)

	for i := length; i > 0; i-- {
		handler = p.middlewares[i-1](handler)
	}

	return handler
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
	time.Sleep(time.Second*2)
	fmt.Println("处理中")
}

func main() {
	pipeline := &Pipeline{}
	send := pipeline.through(func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func(t time.Time) {fmt.Println("请求的执行时间是", time.Since(t))}(time.Now())

			fmt.Println("请求进来了")
			handlerFunc(w, r)
			time.Sleep(time.Second)
			fmt.Println("请求结束了")
		}
	}, func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			defer func(t time.Time) {fmt.Println("请求的执行时间是", time.Since(t))}(time.Now())

			fmt.Println("第二层进来了")
			handlerFunc(writer, request)
			time.Sleep(time.Second)
			fmt.Println("第二层出去了")
		}
	}).send(Handler)

	http.HandleFunc("/duc", send)
	e := http.ListenAndServe(":8888", nil)
	if e != nil {
		return
	}
}
