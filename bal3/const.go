package bal3

const (
	// radix
	base = 3

	prefix = "0t"
)

//------------------------------------------------------------------------------

const (
	tritMin = -1
	tritMax = +1
)

// {-1, 0, +1}

const (
	tv_T = -1
	tv_0 = 0
	tv_1 = +1
)

// const (
// 	tritNegative = -1
// 	tritZero     = 0
// 	tritPositive = +1
// )

var tritsAll = [...]Trit{
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
// 	tc_0 = '0'
// 	tc_1 = '1'
// )

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
