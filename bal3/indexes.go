package bal3

// tritToIndex:
// -1 -> 0
//  0 -> 1
// +1 -> 2

func tritToIndex(t Trit) (index int) {
	index = int(t + 1)
	return index
}

// indexToTrit:
// 0 -> -1
// 1 ->  0
// 2 -> +1

func indexToTrit(index int) (t Trit) {
	t = Trit(index - 1)
	return t
}
