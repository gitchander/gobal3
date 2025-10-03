package ternary

import (
	"fmt"
	"testing"
)

type sampleUnaryFunc struct {
	name     string
	testFunc UnaryFunc
}

type sampleBinaryFunc struct {
	name     string
	testFunc BinaryFunc
}

// wantFunc - reference function
func testUnaryFunc(vs []sampleUnaryFunc, wantFunc UnaryFunc) error {
	n := len(tritValues)
	for _, v := range vs {
		for i := 0; i < n; i++ {
			a := tritValues[i]
			var (
				have = v.testFunc(a)
				want = wantFunc(a)
			)
			if have != want {
				return fmt.Errorf("invalid func %s(%d) result: have %d, want %d", v.name, a, have, want)
			}
		}
	}
	return nil
}

// wantFunc - reference function
func testBinaryFunc(vs []sampleBinaryFunc, wantFunc BinaryFunc) error {
	n := len(tritValues)
	for _, v := range vs {
		for i := 0; i < n; i++ {
			a := tritValues[i]
			for j := 0; j < n; j++ {
				b := tritValues[j]
				var (
					have = v.testFunc(a, b)
					want = wantFunc(a, b)
				)
				if have != want {
					return fmt.Errorf("invalid func %s(%d, %d) result: have %d, want %d", v.name, a, b, have, want)
				}
			}
		}
	}
	return nil
}

func TestInverse(t *testing.T) {

	vs := []sampleUnaryFunc{
		{
			name:     "base-inverse",
			testFunc: BaseCore{}.Inverse,
		},
		{
			name:     "amin-inverse",
			testFunc: AminCore{}.Inverse,
		},
		{
			name:     "amax-inverse",
			testFunc: AmaxCore{}.Inverse,
		},
	}

	wantFunc := func(a Tri) Tri {
		return -a
	}

	err := testUnaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMin(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "base-min",
			testFunc: BaseCore{}.Min,
		},
		{
			name:     "amin-min",
			testFunc: AminCore{}.Min,
		},
		{
			name:     "amax-min",
			testFunc: AmaxCore{}.Min,
		},
	}

	wantFunc := func(a, b Tri) Tri {
		return Min2(a, b)
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMax(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "base-max",
			testFunc: BaseCore{}.Max,
		},
		{
			name:     "amin-max",
			testFunc: AminCore{}.Max,
		},
		{
			name:     "amax-max",
			testFunc: AmaxCore{}.Max,
		},
	}

	wantFunc := func(a, b Tri) Tri {
		return Max2(a, b)
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestXmax(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "base-xmax",
			testFunc: BaseCore{}.Xmax,
		},
		{
			name:     "amin-xmax",
			testFunc: AminCore{}.Xmax,
		},
		{
			name:     "amax-xmax",
			testFunc: AmaxCore{}.Xmax,
		},
	}

	wantFunc := func(a, b Tri) Tri {
		return -(a * b)
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestXamax(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "base_xamax",
			testFunc: BaseCore{}.Xamax,
		},
		{
			name:     "amin_xamax",
			testFunc: AminCore{}.Xamax,
		},
		{
			name:     "amax_xamax",
			testFunc: AmaxCore{}.Xamax,
		},
	}

	wantFunc := func(a, b Tri) Tri {
		return a * b
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}
