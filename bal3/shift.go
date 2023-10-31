package bal3

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

func shiftLeftV1[T Unsigned](tc TryteCore[T], a T, i int) T {
	var (
		mask   = T(tryteMaskTable[tc.n-1])
		offset = i * bitsPerTrit
	)
	return (a << offset) & mask
}

func shiftRightV1[T Unsigned](tc TryteCore[T], a T, i int) T {
	var (
		mask   = T(tryteMaskTable[tc.n-1])
		offset = i * bitsPerTrit
	)
	return (a & mask) >> offset
}

//------------------------------------------------------------------------------

func shiftLeftV2[T Unsigned](tc TryteCore[T], a T, i int) T {
	checkShiftAmount(i)
	var b T
	for j := i; j < tc.n; j++ {
		b = setTrit(b, j, getTrit(a, j-i))
	}
	return b
}

func shiftRightV2[T Unsigned](tc TryteCore[T], a T, i int) T {
	checkShiftAmount(i)
	var b T
	for j := tc.n - 1 - i; j >= 0; j-- {
		b = setTrit(b, j, getTrit(a, j+i))
	}
	return b
}

//------------------------------------------------------------------------------

// Shl - shift left
// a << i
func shiftLeft[T Unsigned](tc TryteCore[T], a T, i int) T {
	//return shiftLeftV1(tc, a, i)
	return shiftLeftV2(tc, a, i)
}

// Shr - shift right
// a >> i
func shiftRight[T Unsigned](tc TryteCore[T], a T, i int) T {
	//return shiftRightV1(tc, a, i)
	return shiftRightV2(tc, a, i)
}

//------------------------------------------------------------------------------
