package bal3

import (
	"errors"
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

func quoRemInt(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

// min <= value < max
func valueIn(value int, min, max int) bool {
	return (min <= value) && (value < max)
}

func not(b bool) bool {
	return !b
}

// min < 0 < max
func bunchesN(x int, min, max int, ds []int) {

	// radix
	base := max - min + 1
	var rem int

	for i := range ds {
		switch {
		case x == 0:
			ds[i] = 0
		case x > 0:
			x, rem = quoRemInt(x-min, base)
			ds[i] = (rem + min)
		case x < 0:
			x, rem = quoRemInt(x-max, base)
			ds[i] = (rem + max)
		}
	}
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
