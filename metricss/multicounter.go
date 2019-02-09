package metrics

import (
)

type MultiCounter struct {
	Name string
	Counters map[string]*Counter
	t uint64
}

func (self *MultiCounter) Reset(counter string) {
	
	if self.Counters[counter]== nil {
		self.Counters[counter] = &Counter{
			Name : counter,
		}
	}

	self.Counters[counter].Reset()
}

func (self *MultiCounter) Increment(counter string) {
	
	if _, ok := self.Counters[counter]; !ok {
		self.Counters[counter] = &Counter{
			Name : counter,
		}
	}	
		
	self.Counters[counter].Increment()
}
