package metrics

import (
)

type Counter struct {
	Name string
	t uint64
}

func (self *Counter) GetValue() uint64 {
	return self.t
}

func (self *Counter) Reset() {
	self.t = 0
}

func (self *Counter) Increment() {
	self.t += 1
}

func (self *Counter) Update(results *map[string]interface{}) error {
	res := *results
	
	res[self.Name] = self.t

	return nil
}