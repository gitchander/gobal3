package bools

import (
	"fmt"
	"testing"
)

//------------------------------------------------------------------------------

// Unary operation
// https://en.wikipedia.org/wiki/Unary_operation

type UnaryFunc func(bool) bool

// Binary operation
// https://en.wikipedia.org/wiki/Binary_operation

type BinaryFunc func(bool, bool) bool

//------------------------------------------------------------------------------

type sampleUnaryFunc struct {
	arg bool
	res bool
}

type sampleBinaryFunc struct {
	arg1, arg2 bool
	res        bool
}

//------------------------------------------------------------------------------

func testUfSample(name string, uf UnaryFunc, sample sampleUnaryFunc) error {
	var (
		have = uf(sample.arg)
		want = sample.res
	)
	if have != want {
		return fmt.Errorf("invalid %q: have %t, want %t", name, have, want)
	}
	return nil
}

func testBfSample(name string, bf BinaryFunc, sample sampleBinaryFunc) error {
	var (
		have = bf(sample.arg1, sample.arg2)
		want = sample.res
	)
	if have != want {
		return fmt.Errorf("invalid %q: have %t, want %t", name, have, want)
	}
	return nil
}

//------------------------------------------------------------------------------

func testUfSamples(name string, uf UnaryFunc, samples []sampleUnaryFunc) error {
	for _, sample := range samples {
		err := testUfSample(name, uf, sample)
		if err != nil {
			return err
		}
	}
	return nil
}

func testBfSamples(name string, bf BinaryFunc, samples []sampleBinaryFunc) error {
	for _, sample := range samples {
		err := testBfSample(name, bf, sample)
		if err != nil {
			return err
		}
	}
	return nil
}

//------------------------------------------------------------------------------

func testCore(name string, c Core) error {

	// Test NOT
	{
		var (
			fname   = fmt.Sprintf("%s.%s", name, "not")
			uf      = c.Not
			samples = []sampleUnaryFunc{
				{arg: true, res: false},
				{arg: false, res: true},
			}
		)
		err := testUfSamples(fname, uf, samples)
		if err != nil {
			return err
		}
	}

	// Test OR
	{
		var (
			fname   = fmt.Sprintf("%s.%s", name, "or")
			bf      = c.Or
			samples = []sampleBinaryFunc{
				{arg1: false, arg2: false, res: false},
				{arg1: true, arg2: false, res: true},
				{arg1: false, arg2: true, res: true},
				{arg1: true, arg2: true, res: true},
			}
		)
		err := testBfSamples(fname, bf, samples)
		if err != nil {
			return err
		}
	}

	// Test AND
	{
		var (
			fname   = fmt.Sprintf("%s.%s", name, "and")
			bf      = c.And
			samples = []sampleBinaryFunc{
				{arg1: false, arg2: false, res: false},
				{arg1: true, arg2: false, res: false},
				{arg1: false, arg2: true, res: false},
				{arg1: true, arg2: true, res: true},
			}
		)
		err := testBfSamples(fname, bf, samples)
		if err != nil {
			return err
		}
	}

	// Test NOR
	{
		var (
			fname   = fmt.Sprintf("%s.%s", name, "nor")
			bf      = c.Nor
			samples = []sampleBinaryFunc{
				{arg1: false, arg2: false, res: true},
				{arg1: true, arg2: false, res: false},
				{arg1: false, arg2: true, res: false},
				{arg1: true, arg2: true, res: false},
			}
		)
		err := testBfSamples(fname, bf, samples)
		if err != nil {
			return err
		}
	}

	// Test NAND
	{
		var (
			fname   = fmt.Sprintf("%s.%s", name, "nand")
			bf      = c.Nand
			samples = []sampleBinaryFunc{
				{arg1: false, arg2: false, res: true},
				{arg1: true, arg2: false, res: true},
				{arg1: false, arg2: true, res: true},
				{arg1: true, arg2: true, res: false},
			}
		)
		err := testBfSamples(fname, bf, samples)
		if err != nil {
			return err
		}
	}

	// Test XOR
	{
		var (
			fname   = fmt.Sprintf("%s.%s", name, "xor")
			bf      = c.Xor
			samples = []sampleBinaryFunc{
				{arg1: false, arg2: false, res: false},
				{arg1: true, arg2: false, res: true},
				{arg1: false, arg2: true, res: true},
				{arg1: true, arg2: true, res: false},
			}
		)
		err := testBfSamples(fname, bf, samples)
		if err != nil {
			return err
		}
	}

	// Test XNOR
	{
		var (
			fname   = fmt.Sprintf("%s.%s", name, "xnor")
			bf      = c.Xnor
			samples = []sampleBinaryFunc{
				{arg1: false, arg2: false, res: true},
				{arg1: true, arg2: false, res: false},
				{arg1: false, arg2: true, res: false},
				{arg1: true, arg2: true, res: true},
			}
		)
		err := testBfSamples(fname, bf, samples)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestBaseCore(t *testing.T) {
	var c Core = BaseCore{}
	err := testCore("base_core", c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNandCore(t *testing.T) {
	var c Core = NandCore{}
	err := testCore("nand_core", c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNorCore(t *testing.T) {
	var c Core = NorCore{}
	err := testCore("nor_core", c)
	if err != nil {
		t.Fatal(err)
	}
}
