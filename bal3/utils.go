package bal3

import (
	"errors"
	"fmt"
)

var powersOfThree = [...]int{
	0:  1,
	1:  3,
	2:  9,
	3:  27,
	4:  81,
	5:  243,
	6:  729,
	7:  2187,
	8:  6561,
	9:  19683,
	10: 59049,
	11: 177147,
	12: 531441,
	13: 1594323,
	14: 4782969,
	15: 14348907,
	16: 43046721,
	17: 129140163,
	18: 387420489,
	19: 1162261467,
	20: 3486784401,
	21: 10460353203,
	22: 31381059609,
	23: 94143178827,
	24: 282429536481,
	25: 847288609443,
	26: 2541865828329,
	27: 7625597484987,
	28: 22876792454961,
	29: 68630377364883,
	30: 205891132094649,
	31: 617673396283947,
	32: 1853020188851841,
	33: 5559060566555523,
	34: 16677181699666569,
	35: 50031545098999707,
	36: 150094635296999121,
	37: 450283905890997363,
	38: 1350851717672992089,
	39: 4052555153018976267,
}

var errNegativeShift = errors.New("negative shift amount")

func checkShiftAmount(i int) {
	if i < 0 {
		panic(errNegativeShift)
	}
}

// TryteMinMax
// n - number of trits.
func tryteBounds(n int) (min, max int) {
	max = (powersOfThree[n] - 1) / 2
	min = -max
	return
}

// a^n
func powN(a int, n int) int {
	p := 1
	for i := 0; i < n; i++ {
		p *= a
	}
	return p
}

// inInterval returns true if (min <= a < max).
func inInterval(a int, min, max int) bool {
	// Check empty interval
	if min >= max {
		return false
	}
	return (min <= a) && (a < max)
}

func not(b bool) bool {
	return !b
}

const bitsPerByte = 8

func bitsPerUnsigned[T Unsigned]() int {
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

// sum 4 trits: [-4..4]
func splitTrits2(v int) (t1, t0 int) {
	switch v {
	case -4:
		return tv_T, tv_T
	case -3:
		return tv_T, tv_0
	case -2:
		return tv_T, tv_1
	case -1:
		return tv_0, tv_T
	case 0:
		return tv_0, tv_0
	case 1:
		return tv_0, tv_1
	case 2:
		return tv_1, tv_T
	case 3:
		return tv_1, tv_0
	case 4:
		return tv_1, tv_1
	default:
		panic(fmt.Errorf("splitTrits2: invalid value %d", v))
	}
}

//------------------------------------------------------------------------------

func parseTable(sss ...string) ([][]int, error) {
	cols := 0
	for _, ss := range sss {
		cols = maxInt(cols, len(ss))
	}
	ttt := make([][]int, len(sss))
	for i, ss := range sss {
		var (
			chars = []byte(ss)
			tt    = make([]int, cols)
		)
		for j, char := range chars {
			t, ok := charToTrit(char)
			if !ok {
				return nil, fmt.Errorf("invalid trit char %q", char)
			}
			tt[j] = t
		}
		ttt[i] = tt
	}
	return ttt, nil
}

func mustParseTable(sss ...string) [][]int {
	table, err := parseTable(sss...)
	if err != nil {
		panic(err)
	}
	return table
}

func tritByTable(table [][]int, a, b int) int {
	return table[a+1][b+1]
}

//------------------------------------------------------------------------------

func quoRemMinMax(x int, min, max int) (q, r int) {

	if min > max {
		panic("min > max")
	}

	// radix
	base := max - min + 1

	if (min <= x) && (x <= max) {
		q, r = 0, x
	}

	if x < min {
		q, r = quoRem(x-max, base)
		r += max
	}

	if x > max {
		q, r = quoRem(x-min, base)
		r += min
	}

	return q, r
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}
