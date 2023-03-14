package counter

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func CounterWithAtomic() {
	var (
		v  uint64
		wg sync.WaitGroup
	)

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&v, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("v:", v)
}
