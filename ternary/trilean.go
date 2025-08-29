package ternary

import (
	"fmt"
)

type Tri int

var (
	tritValues = [...]Tri{-1, 0, +1}

	tritCharsV1 = [...]byte{'T', '0', '1'}
	tritCharsV2 = [...]byte{'-', '0', '+'}
	tritCharsV3 = [...]byte{'N', '0', 'P'}
	tritCharsV4 = [...]byte{'N', '0', '1'}

	tritChars = tritCharsV1
)

func tritToChar(t Tri) byte {
	return tritChars[t+1]
}

func errInvalidTrit(t Tri) error {
	return fmt.Errorf("invalid trit value %d", t)
}

// func checkTrit(t Tri) {
// 	switch t {
// 	case -1, 0, 1:
// 	default:
// 		panic(errInvalidTrit(t))
// 	}
// }

// func checkTrits(ts ...Tri) {
// 	for _, t := range ts {
// 		checkTrit(t)
// 	}
// }

//------------------------------------------------------------------------------

func Min2(a, b Tri) Tri {
	if a < b {
		return a
	}
	return b
}

func Max2(a, b Tri) Tri {
	if a > b {
		return a
	}
	return b
}

func Min3(a, b, c Tri) Tri {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func Max3(a, b, c Tri) Tri {
	if a > b {
		if a > c {
			return a
		}
	} else {
		if b > c {
			return b
		}
	}
	return c
}

func MaxN(as ...Tri) Tri {
	n := len(as)
	if n == 0 {
		return 0 // default value
	}
	j := 0
	for i := 1; i < n; i++ {
		if as[j] < as[i] {
			j = i
		}
	}
	return as[j]
}
