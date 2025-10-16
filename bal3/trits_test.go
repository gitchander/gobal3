package bal3

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
	for _, v := range vs {
		for _, a := range allTrits {
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
	n := len(allTrits)
	for _, v := range vs {
		for i := 0; i < n; i++ {
			a := allTrits[i]
			for j := 0; j < n; j++ {
				b := allTrits[j]
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

func testSplitTrits2(v int) (hi, lo Trit) {
	const (
		vN = -1
		vZ = 0
		vP = +1
	)
	switch v {
	case -2:
		return vN, vP
	case -1:
		return vZ, vN
	case 0:
		return vZ, vZ
	case +1:
		return vZ, vP
	case +2:
		return vP, vN
	default:
		panic(fmt.Errorf("testSplitTrits2: invalid value %d", v))
	}
}

func TestTritsAddSum(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "tritsAddSum",
			testFunc: tritsAddSum,
		},
		{
			name:     "tritsAddSumV1",
			testFunc: tritsAddSumV1,
		},
		{
			name:     "tritsAddSumV2",
			testFunc: tritsAddSumV2,
		},
	}

	wantFunc := func(a, b Trit) Trit {
		_, t0 := testSplitTrits2(int(a + b))
		return t0
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTritsAddCons(t *testing.T) {

	vs := []sampleBinaryFunc{
		{
			name:     "tritsAddCons",
			testFunc: tritsAddCons,
		},
		{
			name:     "tritsAddConsV1",
			testFunc: tritsAddConsV1,
		},
		{
			name:     "tritsAddConsV2",
			testFunc: tritsAddConsV2,
		},
		{
			name:     "tritsAddConsV3",
			testFunc: tritsAddConsV3,
		},
		{
			name:     "tritsAddConsV4",
			testFunc: tritsAddConsV4,
		},
	}

	wantFunc := func(a, b Trit) Trit {
		t1, _ := testSplitTrits2(int(a + b))
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

	wantFunc := func(a, b Trit) Trit {
		return a * b
	}

	err := testBinaryFunc(vs, wantFunc)
	if err != nil {
		t.Fatal(err)
	}
}
