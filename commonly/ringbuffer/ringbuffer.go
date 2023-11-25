package ringbuffer

type ring struct {
	list []interface{}
	size uint
	head uint
	tail uint
}

const defaultSize = 1024

func NewRingBuffer() *ring {
	return &ring{
		list: make([]interface{}, defaultSize),
		size: defaultSize,
	}
}

func (r *ring) mask(index uint) uint {
	return index & (r.size - 1)
}

func (r *ring) Push(data interface{}) {
	if r.full() {
		panic("ring buffer is full")
	}
	r.head++
	r.list[r.mask(r.head)] = data
}

func (r *ring) Pop() interface{} {
	if r.empty() {
		panic("ring buffer is empty")
	}
	r.tail++
	return r.list[r.mask(r.tail)]
}

func (r *ring) full() bool {
	return r.head-r.tail == r.size
}

func (r *ring) empty() bool {
	return r.tail == r.head
}
