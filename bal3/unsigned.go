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
	// Unsigned = customUnsigned
	// Unsigned = constraints.Unsigned
	Signed = constraints.Signed

	coreTryte = constraints.Unsigned
)
