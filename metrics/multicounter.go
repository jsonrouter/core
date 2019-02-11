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

func (self *MultiCounter) Reset(counter string) {
	self.Lock()
	defer self.Unlock()

	if self.Counters[counter]== nil {
		self.Counters[counter] = &Counter{
			Name : counter,
		}
	}

	self.Counters[counter].Reset()
}

func (self *MultiCounter) Increment(counter string) {
	self.Lock()
	defer self.Unlock()

	if _, ok := self.Counters[counter]; !ok {
		self.Counters[counter] = &Counter{
			Name : counter,
		}
	}

	self.Counters[counter].Increment()
}
