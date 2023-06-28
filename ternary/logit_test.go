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

func TestNeg(t *testing.T) {

	vs := []sampleUnaryFunc{
		{
			name:     "Neg",
			testFunc: neg,
		},
		{
			name:     "amin_neg",
			testFunc: aminCore{}.Neg,
		},
		{
			name:     "amax_neg",
			testFunc: amaxCore{}.Neg,
		},
	}

	wantFunc := func(a int) int {
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
			name:     "Min",
			testFunc: min,
		},
		{
			name:     "amin_min",
			testFunc: aminCore{}.Min,
		},
		{
			name:     "amax_min",
			testFunc: amaxCore{}.Min,
		},
	}

	wantFunc := func(a, b int) int {
		return minInt(a, b)
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMax(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "Max",
			testFunc: max,
		},
		{
			name:     "amin_max",
			testFunc: aminCore{}.Max,
		},
		{
			name:     "amax_max",
			testFunc: amaxCore{}.Max,
		},
	}

	wantFunc := func(a, b int) int {
		return maxInt(a, b)
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestXor(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "Xor",
			testFunc: Xor,
		},
		{
			name:     "amin_xor",
			testFunc: aminCore{}.Xor,
		},
		{
			name:     "amax_xor",
			testFunc: amaxCore{}.Xor,
		},
	}

	wantFunc := func(a, b int) int {
		return -(a * b)
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNotXor(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "NotXor",
			testFunc: NotXor,
		},
		{
			name:     "amin_not_xor",
			testFunc: aminCore{}.NotXor,
		},
		{
			name:     "amax_not_xor",
			testFunc: amaxCore{}.NotXor,
		},
	}

	wantFunc := func(a, b int) int {
		return a * b
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCons(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "cons",
			testFunc: cons,
		},
	}

	wantFunc := func(a, b int) int {
		if (a == -1) && (b == -1) {
			return -1
		}
		if (a == 1) && (b == 1) {
			return 1
		}
		return 0
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}
