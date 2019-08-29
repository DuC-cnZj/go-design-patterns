package pool

type poolChan chan Pool

type Pool interface {
	GetNum() int
}

type MyPool struct {
	id int
}

func (m *MyPool) GetNum() int {
	return m.id
}

func New(num int) (p poolChan) {
	p = make(poolChan, num)

	for i := 1; i <= num; i++ {
		p <- &MyPool{id: i}
	}

	return
}
