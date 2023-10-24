package random

import (
	"crypto/rand"
	"encoding/binary"
)

var nextSeed = func() func() int64 {
	const bytesPerUint64 = 8
	seeds := make(chan int64)
	go func() {
		data := make([]byte, bytesPerUint64)
		for {
			_, err := rand.Read(data)
			if err != nil {
				panic(err)
			}
			u := binary.BigEndian.Uint64(data)
			u = (u << 1) >> 1 // clear sign bit
			seeds <- int64(u)
		}
	}()
	return func() int64 {
		return <-seeds
	}
}()
