package iterate

// forward iterate
func ForwardFunc(n int, f func(i int) bool) {
	for i := 0; i < n; i++ {
		if !f(i) {
			return
		}
	}
}

// backward iterate
func BackwardFunc(n int, f func(i int) bool) {
	for i := n; i > 0; {
		i--
		if !f(i) {
			return
		}
	}
}

//------------------------------------------------------------------------------

type Forward struct {
	n int
	i int // index
}

func NewForward(n int) *Forward {
	return &Forward{
		n: n,
		i: -1,
	}
}

func (b *Forward) Next() bool {
	b.i++
	return b.i < b.n
}

func (b *Forward) Index() int {
	return b.i
}

//------------------------------------------------------------------------------

type Backward struct {
	i int // index
}

func NewBackward(n int) *Backward {
	return &Backward{
		i: n,
	}
}

func (b *Backward) Next() bool {
	ok := b.i > 0
	b.i--
	return ok
}

func (b *Backward) Index() int {
	return b.i
}
