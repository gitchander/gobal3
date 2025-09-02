package bal3

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
