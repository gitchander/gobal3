package bal3

import (
	"fmt"
	"math/big"
)

type BigInt = big.Int

func newBigInt(x int64) *BigInt {
	return big.NewInt(x)
}

func bigIntToInt64(b *BigInt) (int64, error) {
	if !(b.IsInt64()) {
		return 0, fmt.Errorf("big int value (%v) is not int64", b)
	}
	return b.Int64(), nil
}

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

func trytesToBigInt[T GenericTryte](n int, as []T, rest *BigInt) *BigInt {
	v := rest
	for i := range backward(len(as)) {
		v = tryteToBigInt(n, as[i], v)
	}
	return v
}
