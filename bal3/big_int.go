package bal3

import (
	"math/big"
)

type BigInt = big.Int

var (
	bigBase = big.NewInt(3)

	bigAllTrits = [3]*big.Int{
		big.NewInt(-1),
		big.NewInt(0),
		big.NewInt(+1),
	}
)

func tryteToBigInt[T GenericTryte](n int, a T, rest *BigInt) *BigInt {

	v := rest
	if v == nil {
		v = big.NewInt(0)
	}

	for i := n; i > 0; { // backward iterate
		i--

		// v = (v * base) + int64(getTrit(a, i))

		index := tritToIndex(getTrit(a, i))
		bt := bigAllTrits[index]

		v.Mul(v, bigBase)
		v.Add(v, bt)
	}
	return v
}

func bigTryteToBigInt(b *BigTryte, rest *BigInt) *BigInt {

	v := rest
	if v == nil {
		v = big.NewInt(0)
	}

	b.backward(func(i int, t Trit) bool {

		index := tritToIndex(t)
		bt := bigAllTrits[index]

		v.Mul(v, bigBase)
		v.Add(v, bt)

		return true
	})

	return v
}
