package metrics

import (
	"sync"
)

type MetricsInterface struct {
	Update error
}

type Metrics struct {
	Timers map[string]*Timer
	Counters map[string]*Counter
	MultiCounters map[string]*MultiCounter
	//Config *config
	Results map[string]interface{}
	sync.RWMutex
}

func (self *MultiCounter) Update(results *map[string]interface{}) error {
	self.Lock()
	defer self.Unlock()

	res := *results

	r := make(map[string]interface{})

	for name, counter := range self.Counters {
		//n := self.Name + ":" + name
		//res[n] = counter.t
		r[name] = counter.t
	}

	res[self.Name] = r

	return nil
}
