package bal3

import (
	"github.com/gitchander/gobal3/ternary"
)

type (
	UnaryFunc  func(Trit) Trit
	BinaryFunc func(Trit, Trit) Trit
)

type triCore struct{}

func (triCore) Neg(a Trit) Trit {
	c := ternary.Neg(ternary.Tri(a))
	return Trit(c)
}

func (triCore) Min(a, b Trit) Trit {
	c := ternary.Min(ternary.Tri(a), ternary.Tri(b))
	return Trit(c)
}

func (triCore) Max(a, b Trit) Trit {
	c := ternary.Max(ternary.Tri(a), ternary.Tri(b))
	return Trit(c)
}

func (triCore) Is(a Trit, v Trit) Trit {
	c := ternary.Is(ternary.Tri(a), ternary.Tri(v))
	return Trit(c)
}

func (triCore) Inc(a Trit) Trit {
	c := ternary.Inc(ternary.Tri(a))
	return Trit(c)
}

func (triCore) Dec(a Trit) Trit {
	c := ternary.Dec(ternary.Tri(a))
	return Trit(c)
}

func (triCore) Xmax(a, b Trit) Trit {
	c := ternary.Xmax(ternary.Tri(a), ternary.Tri(b))
	return Trit(c)
}

func (triCore) Xamax(a, b Trit) Trit {
	c := ternary.Xamax(ternary.Tri(a), ternary.Tri(b))
	return Trit(c)
}

var trico triCore
