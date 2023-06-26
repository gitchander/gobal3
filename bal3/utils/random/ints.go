package random

func RandInts(r *Rand, n int, min, max int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = min + r.Intn(max-min)
	}
	return a
}
