package geom

//------------------------------------------------------------------------------

func not(b bool) bool {
	return !b
}

//------------------------------------------------------------------------------

type Coord = Point2i

func MakeCoord(x, y int) Coord {
	return MakePoint2i(x, y)
}

//------------------------------------------------------------------------------

func clampMin(a, min int) int {
	if a < min {
		a = min
	}
	return a
}

//------------------------------------------------------------------------------

type Matrix2D[T any] struct {
	size   Coord
	values []T
}

func NewMatrix2D[T any](size Coord) *Matrix2D[T] {

	size.X = clampMin(size.X, 0)
	size.Y = clampMin(size.Y, 0)

	return &Matrix2D[T]{
		size:   size,
		values: make([]T, (size.X * size.Y)),
	}
}

func (p *Matrix2D[T]) checkCoord(coord Coord) bool {
	if not((0 <= coord.X) && (coord.X < p.size.X)) {
		return false
	}
	if not((0 <= coord.Y) && (coord.Y < p.size.Y)) {
		return false
	}
	return true
}

func (p *Matrix2D[T]) coordOffset(coord Coord) int {
	return (coord.Y * p.size.X) + coord.X
}

func (p *Matrix2D[T]) Size() Coord {
	return p.size
}

func (p *Matrix2D[T]) GetValue(coord Coord) (T, bool) {
	if not(p.checkCoord(coord)) {
		var zv T
		return zv, false
	}
	var (
		offset = p.coordOffset(coord)
		value  = p.values[offset]
	)
	return value, true
}

func (p *Matrix2D[T]) SetValue(coord Coord, value T) bool {
	if not(p.checkCoord(coord)) {
		return false
	}
	offset := p.coordOffset(coord)
	p.values[offset] = value
	return true
}

func (p *Matrix2D[T]) GetValueXY(x, y int) (T, bool) {
	return p.GetValue(MakeCoord(x, y))
}

func (p *Matrix2D[T]) SetValueXY(x, y int, value T) bool {
	return p.SetValue(MakeCoord(x, y), value)
}

func (p *Matrix2D[T]) Walk(f func(coord Coord, value T) bool) {
	for y := 0; y < p.size.Y; y++ {
		for x := 0; x < p.size.X; x++ {
			var (
				coord  = MakeCoord(x, y)
				offset = p.coordOffset(coord)
				value  = p.values[offset]
			)
			if not(f(coord, value)) {
				return
			}
		}
	}
}

func (p *Matrix2D[T]) Modify(f func(coord Coord, value T) (T, bool)) {
	for y := 0; y < p.size.Y; y++ {
		for x := 0; x < p.size.X; x++ {
			var (
				coord  = MakeCoord(x, y)
				offset = p.coordOffset(coord)
				value  = p.values[offset]
			)
			newValue, ok := f(coord, value)
			if !ok {
				return
			}
			p.values[offset] = newValue
		}
	}
}

//------------------------------------------------------------------------------
