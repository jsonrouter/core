package metrics

import (
	"sync"
)

type BenchMark struct { 
	Name string
	Timers map[string]*Timer
	sync.RWMutex
}

func (self *BenchMark) StartTimer(timer string) {
	t := self.Timers[timer] 
	if t != nil {
		self.Timers[timer].Start()
	}	
}

func (self *BenchMark) StopTimer(timer string) {
	t := self.Timers[timer] 
	if t != nil {
		self.Timers[timer].Stop()
	}
}

func (self *BenchMark) StartTimers() {
	for _, timer := range self.Timers{
		timer.Start()
		//go timer.Start()
	}	
}

func (self *BenchMark) StopTimers() {
	for _, timer := range self.Timers{
		timer.Stop()
		//go timer.Stop(
	}	
}


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

