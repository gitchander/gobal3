package random

import (
	"math/rand"
)

type Rand = rand.Rand

func RandBool(r *Rand) bool {
	return (r.Uint32() & 1) == 1
}

func RandIntMinMax(r *Rand, min, max int) int {
	return min + r.Intn(max-min)
}
