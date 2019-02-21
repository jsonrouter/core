package metrics

import (
	"sync"

)

type Counter struct {
	Name string
	t uint64
	sync.RWMutex
}

// GetValue exposes the value
func (self *Counter) GetValue() uint64 {
	self.RLock()
	defer self.RUnlock()

	return self.t
}

// Reset will reset the counter
func (self *Counter) Reset() {
	self.Lock()
	defer self.Unlock()

	self.t = 0
}

// Increment
func (self *Counter) Increment() {
	self.Lock()
	defer self.Unlock()

	self.t += 1
}

// Update
func (self *Counter) Update(results *map[string]interface{}) error {

	self.Lock()
	defer self.Unlock()

	res := *results
	res[self.Name] = self.t
	
	return nil
}
