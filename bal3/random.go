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
	return random.RandByCorpus(r, tritValues[:])
}
