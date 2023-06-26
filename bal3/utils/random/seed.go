package random

var nextSeed = func() func() int64 {
	c := make(chan int64)
	go func() {
		const period = 1000
		for {
			r := NewRandNow()
			for i := 0; i < period; i++ {
				c <- r.Int63()
			}
		}
	}()
	return func() int64 {
		return <-c
	}
}()
