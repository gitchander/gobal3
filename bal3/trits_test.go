package bal3

import (
	"fmt"
	"testing"

	. "github.com/gitchander/gobal3/ternary"
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
	for _, v := range vs {
		for _, a := range tritValues {
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

func TestAddSum(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "addSum",
			testFunc: addSum,
		},
		{
			name:     "addSumV1",
			testFunc: addSumV1,
		},
		{
			name:     "addSumV2",
			testFunc: addSumV2,
		},
		{
			name:     "addSumV3",
			testFunc: addSumV3,
		},
		{
			name:     "addSumV4",
			testFunc: addSumV4,
		},
	}

	wantFunc := func(a, b int) int {
		_, t0 := splitTrits2(a + b)
		return t0
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddCons(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "addCons",
			testFunc: addCons,
		},
		{
			name:     "addConsV1",
			testFunc: addConsV1,
		},
		{
			name:     "addConsV2",
			testFunc: addConsV2,
		},
		{
			name:     "addConsV3",
			testFunc: addConsV3,
		},
		{
			name:     "addConsV4",
			testFunc: addConsV4,
		},
		{
			name:     "addConsV5",
			testFunc: addConsV5,
		},
	}

	wantFunc := func(a, b int) int {
		t1, _ := splitTrits2(a + b)
		return t1
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTritsMul(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "tritsMul",
			testFunc: tritsMul,
		},
		{
			name:     "tritsMulV1",
			testFunc: tritsMulV1,
		},
		{
			name:     "tritsMulV2",
			testFunc: tritsMulV2,
		},
		{
			name:     "tritsMulV3",
			testFunc: tritsMulV3,
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
