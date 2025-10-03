package ternary

type BaseCore struct{}

var _ Core = BaseCore{}

func (BaseCore) Inverse(a Tri) Tri {
	return Inverse(a)
}

func (BaseCore) Min(a, b Tri) Tri {
	return Min(a, b)
}

func (BaseCore) Max(a, b Tri) Tri {
	return Max(a, b)
}

func (BaseCore) Xmax(a, b Tri) Tri {
	return Xmax(a, b)
}

func (BaseCore) Xamax(a, b Tri) Tri {
	return Xamax(a, b)
}
