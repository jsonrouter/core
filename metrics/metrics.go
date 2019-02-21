// Package metrics is a package providing counters and timers for measuring various processes. Common usages
// might include measuing request counts or average request time.

package metrics

import (
	"sync"
)

type MetricsInterface struct {
	Update error
}

// Metrics is the main object to instantiate when using the metrics package. One instance of
// a Metrics object will contain all your timers and counters and results
type Metrics struct {
	Timers map[string]*Timer
	Counters map[string]*Counter
	MultiCounters map[string]*MultiCounter
	BenchMarks map[string]*BenchMark
	//Config *config
	Results map[string]interface{}
	sync.RWMutex
}

func sum (vals ...uint64) uint64 {

	var sum uint64
	
	for _, val := range vals {
		sum += val
	}
	
	return sum
}
