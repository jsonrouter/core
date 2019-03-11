// Package metrics is a package providing counters and timers for measuring various processes. Common usages
// might include measuing request counts or average request time.

package metrics

import (
	"sync"
	"encoding/json"
)

/*
type MetricsInterface struct {
	Update error
}
*/

func NewMetrics() Metrics {
	return Metrics{
		Timers: map[string]*Timer{
			"requestTime": &Timer{
				Name : "requestTime",
				BufferSize : 1000,
			},
		},
		Counters: map[string]*Counter{
			"requestCount" : &Counter{
				Name : "requestCount",
			},
		},
		MultiCounters: map[string]*MultiCounter{
			"responseCodes" : &MultiCounter{
				Name : "responseCodes",
				Counters : map[string]*Counter{},
			},
			"requestMethods" : &MultiCounter{
				Name : "requestMethods",
				Counters : map[string]*Counter{},
			},
		},
		results: &Results{
			results: map[string]interface{}{},
		},
	}
}

type Results struct {
	results map[string]interface{}
	sync.RWMutex
}

// Metrics is the main object to instantiate when using the metrics package. One instance of
// a Metrics object will contain all your timers and counters and results
type Metrics struct {
	Timers map[string]*Timer
	Counters map[string]*Counter
	MultiCounters map[string]*MultiCounter
	results *Results
	sync.RWMutex
}

func (self *Metrics) MarshalResults() ([]byte, error) {
	self.results.RLock()
	defer self.results.RUnlock()
	return json.Marshal(self.results.results)
}

func (self *Metrics) SetResults(k string, v interface{}) {
	self.results.Lock()
	defer self.results.Unlock()
	self.results.results[k] = v
}

func sum(vals ...uint64) uint64 {

	var sum uint64

	for _, val := range vals {
		sum += val
	}

	return sum
}
