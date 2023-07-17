package overflows

import (
	"math"
)

// https://stackoverflow.com/questions/33641717/detect-signed-int-overflow-in-go

func AddInt64(a, b int64) (int64, bool) {
	const (
		min = math.MinInt64
		max = math.MaxInt64
	)
	if a < 0 {
		if b < (min - a) {
			return 0, false
		}
	}
	if a > 0 {
		if b > (max - a) {
			return 0, false
		}
	}
	return a + b, true
}
