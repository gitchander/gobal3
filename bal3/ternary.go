package bal3

import (
	"github.com/gitchander/gobal3/ternary"
)

type UnaryFunc func(Trit) Trit
type BinaryFunc func(Trit, Trit) Trit

func terNeg(a Trit) Trit {
	c := ternary.Neg(int(a))
	return Trit(c)
}

func terMin(a, b Trit) Trit {
	c := ternary.Min(int(a), int(b))
	return Trit(c)
}

func terMax(a, b Trit) Trit {
	c := ternary.Max(int(a), int(b))
	return Trit(c)
}

func terIs(a Trit, v Trit) Trit {
	c := ternary.Is(int(a), int(v))
	return Trit(c)
}

func terInc(a Trit) Trit {
	c := ternary.Inc(int(a))
	return Trit(c)
}

func terDec(a Trit) Trit {
	c := ternary.Dec(int(a))
	return Trit(c)
}

func terNegXor(a, b Trit) Trit {
	c := ternary.NegXor(int(a), int(b))
	return Trit(c)
}
