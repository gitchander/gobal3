package bal3

// TryteMinMax
// n - number of trits.
// func tryteBounds(n int) (min, max int64) {
// 	max = (powersOfThree[n] - 1) / 2
// 	min = -max
// 	return
// }

// a^n
func powN(a int, n int) int {
	p := 1
	for i := 0; i < n; i++ {
		p *= a
	}
	return p
}

func not(b bool) bool {
	return !b
}

const bitsPerByte = 8

func bitsPerUnsigned[T GenericTryte]() int {
	x := uint64(^T(0))
	count := 0
	for x != 0 {
		x >>= bitsPerByte
		count += bitsPerByte
	}
	return count
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func setAllInts(as []int, v int) {
	for i := range as {
		as[i] = v
	}
}

func ceilDiv(a, b int) int {
	return (a + (b - 1)) / b
}

//------------------------------------------------------------------------------

func quoRemBal3(a int64) (q, r int64) {

	const (
		min = -1
		max = +1

		base = max - min + 1
	)

	q, r = quoRem(a, base)

	if r < min {
		q--
		r += base
	}
	if r > max {
		q++
		r -= base
	}

	return q, r
}

//------------------------------------------------------------------------------

func quoRem[T signed](a, b T) (quo, rem T) {
	quo = a / b
	rem = a % b
	return
}

//------------------------------------------------------------------------------

func reverseSlice[T any](as []T) {
	i, j := 0, (len(as) - 1)
	for i < j {
		as[i], as[j] = as[j], as[i]
		i, j = i+1, j-1
	}
}
