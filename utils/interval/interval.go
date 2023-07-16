package interval

type Interval struct {
	Min, Max int
}

func Ivl(min, max int) Interval {
	return Interval{
		Min: min,
		Max: max,
	}
}

func (x Interval) Empty() bool {
	return x.Min >= x.Max
}

func (x Interval) notEmpty() bool {
	return x.Min < x.Max
}

func (x Interval) Width() int {
	if x.Empty() {
		return 0
	}
	return x.Max - x.Min
}

func (x Interval) Contains(v int) bool {
	if x.Empty() {
		return false
	}
	return (x.Min <= v) && (v < x.Max)
}

func (x Interval) Overlaps(y Interval) bool {
	if x.Empty() || y.Empty() {
		return false
	}
	return (x.Min < y.Max) && (y.Min < x.Max)
}
