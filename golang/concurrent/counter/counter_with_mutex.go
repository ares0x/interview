package counter

import (
	"sync"
	"container/list"
)

type CounterWithMutex struct {
	sync.Mutex
	count int
}

func NewCounterWithMutex() *CounterWithMutex {
	return &CounterWithMutex{}
}

func (c *CounterWithMutex) Incr() {
	c.Lock()
	defer c.Unlock()

	c.count++
}

func (c *CounterWithMutex) Decr() {
	c.Lock()
	defer c.Unlock()
	c.count--
}

func (c *CounterWithMutex) Count() int{
	c.Lock()
	defer c.Unlock()
	return c.count
}
