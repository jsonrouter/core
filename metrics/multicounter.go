package metrics

import (
	"sync"
)

type MultiCounter struct {
	Name string
	Counters map[string]*Counter
	t uint64
	sync.RWMutex
}

func (self *MultiCounter) Get(counter string) (c *Counter) {
	// if counter exists do a read lock and return counter
	self.RLock()
	c = self.Counters[counter]
	self.RUnlock()
	if c != nil {
		return c
	}
	// lock the whole process so it is atomic
	self.Lock()
	c = self.Counters[counter]
	if c == nil {
		// init the counter
		c = &Counter{
			Name : counter,
		}
		self.Counters[counter] = c
	}
	// unlock before we call any other functions
	self.Unlock()

	return c
}

func (self *MultiCounter) Reset(counter string) {
	c := self.Get(counter)
	c.Reset()
}

func (self *MultiCounter) Increment(counter string) {
	c := self.Get(counter)
	c.Increment()
}
