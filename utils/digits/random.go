package digits

import (
	"math/rand"

	"github.com/gitchander/gobal3/utils/random"
)

func randNow() *rand.Rand {
	return random.NewRandNow()
}

func randBySeed(seed int64) *rand.Rand {
	return random.NewRandSeed(seed)
}

func randBool(r *rand.Rand) bool {
	return (r.Int() & 1) == 1
}

func randBaseInt(r *rand.Rand) int {
	var (
		k    = 10
		mask = uint32((1 << k) - 1)
	)
	x := (r.Uint32() & mask) >> r.Intn(k)
	a := int(x)
	if randBool(r) {
		a = -a
	}
	return a
}

func randomBaseMinMax(r *rand.Rand) (min, max int) {
	var a, b int
	for a == b {
		a = randBaseInt(r)
		b = randBaseInt(r)
	}
	if a > b {
		a, b = b, a
	}
	return a, b
}
