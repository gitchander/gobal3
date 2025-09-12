package bal3

import (
	"golang.org/x/exp/constraints"
)

// Unsigned integers
type customUnsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func _[T constraints.Unsigned]() {}

type (
	signed = constraints.Signed

	GenericTryte = constraints.Unsigned
)
