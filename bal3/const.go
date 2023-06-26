package bal3

const (
	// radix
	base = 3

	prefix = "0t"
)

//------------------------------------------------------------------------------

// {-1, 0, +1}

const (
	tv_T = -1
	tv_0 = 0
	tv_1 = +1
)

var tritValues = [...]int{
	tv_T,
	tv_0,
	tv_1,
}

//------------------------------------------------------------------------------

// {T, 0, 1}

const (
	tc_T = 'T'
	tc_0 = '0'
	tc_1 = '1'
)

// const (
// 	tc_T = 'N'
// 	tc_0 = 'Z'
// 	tc_1 = 'P'
// )

var tritChars = [...]byte{
	tc_T,
	tc_0,
	tc_1,
}

//------------------------------------------------------------------------------

// const TritsPerTryte4 = 4
// const TritsPerTryte8 = 8
// const TritsPerTryte16 = 16
// const TritsPerTryte32 = 32

// const TritsPerTryte6 = 6
// const TritsPerTryte9 = 9
