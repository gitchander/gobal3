package optional

type Optional[T any] struct {
	Present bool
	Value   T
}

func MakePresent[T any](value T) Optional[T] {
	return Optional[T]{
		Present: true,
		Value:   value,
	}
}

func (o Optional[T]) GetValue() (T, bool) {
	if o.Present {
		return o.Value, true
	}
	var zeroValue T
	return zeroValue, false
}

func (o *Optional[T]) SetValue(value T) {
	*o = MakePresent(value)
}

func (o *Optional[T]) Reset() {
	*o = Optional[T]{}
}

func (o Optional[T]) If(f func(T)) {
	if o.Present {
		f(o.Value)
	}
}
