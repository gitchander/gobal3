package zigzag

import (
	"math"
)

//------------------------------------------------------------------------------

func EncodeInt32(x int32) uint32 {
	ux := uint32(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return ux
}

func DecodeInt32(ux uint32) int32 {
	x := int32(ux >> 1)
	if (ux & 1) != 0 {
		x = ^x
	}
	return x
}

//------------------------------------------------------------------------------

func EncodeInt64(x int64) uint64 {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return ux
}

func DecodeInt64(ux uint64) int64 {
	x := int64(ux >> 1)
	if (ux & 1) != 0 {
		x = ^x
	}
	return x
}

//------------------------------------------------------------------------------

func EncodeInt(x int) uint {
	if math.MaxInt == math.MaxInt32 {
		return uint(EncodeInt32(int32(x)))
	}
	return uint(EncodeInt64(int64(x)))
}

func DecodeInt(ux uint) int {
	if math.MaxInt == math.MaxInt32 {
		return int(DecodeInt32(uint32(ux)))
	}
	return int(DecodeInt64(uint64(ux)))
}

//------------------------------------------------------------------------------
