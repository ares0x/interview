package counter

import (
	"sync"
)

type CounterWithRWMutex struct {
	sync.RWMutex
	count int
}

func NewCounterWithRWMutex() *CounterWithRWMutex {
	return &CounterWithRWMutex{}
}

func (c *CounterWithRWMutex) Incr() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func (c *CounterWithRWMutex) Decr() {
	c.Lock()
	defer c.Unlock()
	c.count--
}

func (c *CounterWithRWMutex) Count() int{
	c.RLock()
	defer c.Unlock()
	return c.count
}
