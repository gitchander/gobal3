package bal3

import (
	"errors"
)

var errNegativeShift = errors.New("negative shift amount")

func checkShiftAmount(i int) {
	if i < 0 {
		panic(errNegativeShift)
	}
}

func makeTryteMaskTable() []uint64 {
	ms := make([]uint64, 32)
	var m uint64
	for i := range ms {
		m = (m << bitsPerTrit) | tbs_Mask
		ms[i] = m
	}
	return ms
}

var tryteMaskTable = makeTryteMaskTable()

//------------------------------------------------------------------------------

func shiftLeftV1[T CoreTryte](n int, a T, i int) T {
	var (
		mask   = T(tryteMaskTable[n-1])
		offset = i * bitsPerTrit
	)
	return (a << offset) & mask
}

func shiftRightV1[T CoreTryte](n int, a T, i int) T {
	var (
		mask   = T(tryteMaskTable[n-1])
		offset = i * bitsPerTrit
	)
	return (a & mask) >> offset
}

//------------------------------------------------------------------------------

func shiftLeftV2[T CoreTryte](n int, a T, i int) T {
	checkShiftAmount(i)
	var b T
	for j := i; j < n; j++ {
		b = setTrit(b, j, getTrit(a, j-i))
	}
	return b
}

func shiftRightV2[T CoreTryte](n int, a T, i int) T {
	checkShiftAmount(i)
	var b T
	for j := (n - i) - 1; j >= 0; j-- { // backward iterate
		b = setTrit(b, j, getTrit(a, i+j))
	}
	return b
}

//------------------------------------------------------------------------------

// Shl - shift left
// a << i
func tryteShiftLeft[Tryte CoreTryte](n int, a Tryte, i int) Tryte {
	//return shiftLeftV1(n, a, i)
	return shiftLeftV2(n, a, i)
}

// Shr - shift right
// a >> i
func tryteShiftRight[Tryte CoreTryte](n int, a Tryte, i int) Tryte {
	//return shiftRightV1(n, a, i)
	return shiftRightV2(n, a, i)
}

//------------------------------------------------------------------------------
