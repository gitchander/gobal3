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

func bitsPerUnsigned[T coreTryte]() int {
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

func hasDuplicate[T comparable](as []T) bool {
	m := make(map[T]struct{})
	for _, a := range as {
		if _, ok := m[a]; ok {
			return true
		}
		m[a] = struct{}{}
	}
	return false
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

func quoRem[T Signed](a, b T) (quo, rem T) {
	quo = a / b
	rem = a % b
	return
}

//------------------------------------------------------------------------------

func parseTable(sss ...string) ([][]Trit, error) {
	cols := 0
	for _, ss := range sss {
		cols = maxInt(cols, len(ss))
	}
	ttt := make([][]Trit, len(sss))
	for i, ss := range sss {
		var (
			chars = []byte(ss)
			tt    = make([]Trit, cols)
		)
		for j, char := range chars {
			t, err := charToTrit(char)
			if err != nil {
				return nil, err
			}
			tt[j] = t
		}
		ttt[i] = tt
	}
	return ttt, nil
}

func mustParseTable(sss ...string) [][]Trit {
	table, err := parseTable(sss...)
	if err != nil {
		panic(err)
	}
	return table
}

func tritByTable(table [][]Trit, a, b Trit) Trit {

	// trit to index: (trit + 1)
	// (-1 + 1) = 0
	// ( 0 + 1) = 1
	// (+1 + 1) = 2

	return table[a+1][b+1]
}

//------------------------------------------------------------------------------

// forward iterate
func forward(n int, f func(i int) bool) {
	for i := 0; i < n; i++ {
		if !f(i) {
			return
		}
	}
}

// backward iterate
func backward(n int, f func(i int) bool) {
	for i := n; i > 0; {
		i--
		if !f(i) {
			return
		}
	}
}

//------------------------------------------------------------------------------

func reverseSlice[T any](as []T) {
	i, j := 0, (len(as) - 1)
	for i < j {
		as[i], as[j] = as[j], as[i]
		i, j = i+1, j-1
	}
}
