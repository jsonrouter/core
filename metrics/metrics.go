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
