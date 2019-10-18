package pipe

type myFunc func(...interface{})

type middleware func(myFunc) myFunc

type Pipeline struct {
	middlewares []middleware
}

func (p *Pipeline) through(m ...middleware) *Pipeline {
	p.middlewares = append(p.middlewares, m...)

	return p
}

func (p *Pipeline) send(handler myFunc) myFunc {
	length := len(p.middlewares)

	for i := length; i > 0; i-- {
		handler = p.middlewares[i-1](handler)
	}

	return handler
}