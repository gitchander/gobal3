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
	return random.RandByCorpus(r, allTrits[:])
}

func randTryte[Tryte GenericTryte](n int, r *Rand) Tryte {
	var a Tryte
	for i := 0; i < n; i++ {
		a = setTrit(a, i, randTrit(r))
	}
	return a
}

func randTryteSh[Tryte GenericTryte](n int, r *Rand) Tryte {
	a := randTryte[Tryte](n, r)
	return tryteShr(n, a, r.Intn(n))
}
