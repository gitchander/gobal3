package bal3

import (
	"math/rand"

	"github.com/gitchander/gobal3/utils/random"
)

type Rand = rand.Rand

func newRandNext() *rand.Rand {
	return random.NewRandNext()
}

func randIntByCorpus(r *random.Rand, corpus []int) int {
	return corpus[r.Intn(len(corpus))]
}

func randTrit(r *random.Rand) int {
	return randIntByCorpus(r, tritValues[:])
}
