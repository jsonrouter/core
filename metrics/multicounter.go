package metrics

import (
	"sync"
)

// MultiCounter is a collection of counters. Useful for making comparisons between groups of
// similar statistics
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

// Reset resets specific counter
func (self *MultiCounter) Reset(counter string) {
	c := self.Get(counter)
	c.Reset()
}

// Increment increments specific counter by 1
func (self *MultiCounter) Increment(counter string) {
	c := self.Get(counter)
	c.Increment()
}

// Update is called to save the values of all counters into the results map.
// You can pass any map[string]Interface{} to store results including the provide Results
// map on the main Metrics struct
func (self *MultiCounter) Update(mtx *sync.RWMutex, results map[string]interface{}) {

	r := make(map[string]interface{})

	self.RLock()
	for _, counter := range self.Counters {
		//r[name] = counter.t
		r[counter.Name] = counter.t
	}
	self.RUnlock()

	mtx.Lock()
	defer mtx.Unlock()

	results[self.Name] = r
}
