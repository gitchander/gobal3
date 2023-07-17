package main

import (
	"fmt"
	"log"
	"math"

	"github.com/gitchander/gobal3/utils/overflows"
)

func main() {
	// fmt.Println(math.MinInt / 9)
	// fmt.Println(overflows.MulInt(math.MinInt, 9))
	testSamples()
	//testAdd()
}

func testSamples() {
	samples := []struct {
		a, b, c int64
		success bool
	}{
		{a: 0, b: 0, c: 0, success: true},
		{a: 9223372036854776, b: 1000, c: 0, success: false},
		{a: 4611686018427387904, b: 2, c: 0, success: false},
	}
	for _, sample := range samples {
		c, success := overflows.MulInt64(sample.a, sample.b)
		if false {
			if c != sample.c {
				log.Fatalf("invalid %q: have %d, want %d", "c", c, sample.c)
			}
		}
		if success != sample.success {
			log.Fatalf("invalid %q: have %t, want %t", "success", success, sample.success)
		}
	}
}

func testAdd() {
	min := math.MinInt64
	fmt.Println(min - min)
	fmt.Println(overflows.AddInt64(math.MaxInt, 1))
}
