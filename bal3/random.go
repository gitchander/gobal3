package bal3

import (
	"math/rand"

	"github.com/gitchander/gobal3/utils/random"
)

type Rand = rand.Rand

func newRandNext() *rand.Rand {
	return random.NewRandNext()
}

func randTrit(r *random.Rand) Trit {
	return random.RandByCorpus(r, tritsAll[:])
}

func randTryte[T coreTryte](n int, r *Rand) T {
	var a T
	for i := 0; i < n; i++ {
		a = setTrit(a, i, randTrit(r))
	}
	return a
}

func randTryteSh[T coreTryte](n int, r *Rand) T {
	a := randTryte[T](n, r)
	return tryteShiftRight(n, a, r.Intn(n))
}
