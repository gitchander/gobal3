package optional

//------------------------------------------------------------------------------

type OptInt = Optional[int]

func MakePresentInt(value int) OptInt {
	return MakePresent(value)
}

//------------------------------------------------------------------------------

type OptInt64 = Optional[int64]

func MakePresentInt64(value int64) OptInt64 {
	return MakePresent(value)
}

//------------------------------------------------------------------------------

type OptFloat64 = Optional[float64]

func MakePresentFloat64(value float64) OptFloat64 {
	return MakePresent(value)
}

//------------------------------------------------------------------------------

type OptString = Optional[string]

func MakePresentString(value string) OptString {
	return MakePresent(value)
}

//------------------------------------------------------------------------------
