package bal3

const (
	bitsPerWord  = 64
	tritsPerWord = bitsPerWord / bitsPerTrit
)

type word uint64

var tcWord = MakeTryteCore[Tryte4](4)
