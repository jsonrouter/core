package metrics

import (
	"sync"
)

// BenchMark is an object consisting of a group of timers that can started and stopped one-at-a-time
// or independently. Useful for making comparisons such as average time taken for POST and GET requests
type BenchMark struct { 
	Name string
	Timers map[string]*Timer
	sync.RWMutex
}

// StartTimer starts specific timer
func (self *BenchMark) StartTimer(timer string) {
	t := self.Timers[timer] 
	if t != nil {
		self.Timers[timer].Start()
	}	
}

// StopTimer stops specific timer
func (self *BenchMark) StopTimer(timer string) {
	t := self.Timers[timer] 
	if t != nil {
		self.Timers[timer].Stop()
	}
}

// StartTimers starts all timers
func (self *BenchMark) StartTimers() {
	for _, timer := range self.Timers{
		timer.Start()
		//go timer.Start()
	}	
}

// StopTimer stops all timers
func (self *BenchMark) StopTimers() {
	for _, timer := range self.Timers{
		timer.Stop()
		//go timer.Stop(
	}	
}

// Update is called to save the values of all timers into the results map. 
// You can pass any map[string]Interface{} to store results including the provide Results
// map on the main Metrics struct
func (self *BenchMark) Update(results *map[string]interface{}) error {
	self.Lock()
	defer self.Unlock()

	res := *results

	r := make(map[string]interface{})

	for _, timer := range self.Timers {
		timer.Update(&r)
	}

	res[self.Name] = r

	return nil
}

