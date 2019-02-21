package metrics

import (
	"sync"

)

// Counter is a simple counter
type Counter struct {
	Name string
	t uint64
	sync.RWMutex
}

// GetValue is a method to directly get the current value of the counter
func (self *Counter) GetValue() uint64 {
	self.RLock()
	defer self.RUnlock()

	return self.t
}

// Reset resets the counter to 0
func (self *Counter) Reset() {
	self.Lock()
	defer self.Unlock()

	self.t = 0
}

// Incrememt increments the counter by 1.
// Call this function where the action want to count is.
func (self *Counter) Increment() {
	self.Lock()
	defer self.Unlock()

	self.t += 1
}

// Update is called to save the value the counterinto the results map.
// You can pass any map[string]Interface{} to store results including the provide Results
// map on the main Metrics struct
func (self *Counter) Update(results *map[string]interface{}) error {

	self.Lock()
	defer self.Unlock()

	res := *results
	res[self.Name] = self.t

	return nil
}
