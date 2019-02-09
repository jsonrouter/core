package metrics

import (

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
}

func (self *MultiCounter) Update(results *map[string]interface{}) error {
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
