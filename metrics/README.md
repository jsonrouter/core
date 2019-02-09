# metrics
Metrics

## Usage
  
```
import "metrics"

var met metrics.Metrics{
			Timers: map[string]*metrics.Timer{
				"timerName": &metrics.Timer{
					Name : "timerName",
					BufferSize: 1000,
				},
			},
			Counters: map[string]*metrics.Counter{
				"counterName" : &metrics.Counter{
					Name : "counterName",
				},
			},
			MultiCounters: map[string]*metrics.MultiCounter{
				"multiCounterName" : &metrics.MultiCounter{
					Name : "multiCounterName",
					Counters : map[string]*metrics.Counter{},
				},
			}

var resultsMap map[string]interface{}
```
```
func timerExampele() {
	met.Timers["timerName"].Start()

	//stuff to time

	met.Timers["timerName"].Stop()
	met.Timers["timerName"].Update(&resultsMap)
}

```
```
func counterExample() {

	//stuff to count

	met.Counters["counterName"].Increment()
	met.Counters["counterName"].Update(&resultsMap)
}
```

```
func multiCounterExample() {

	//stuff to count

	met.MultiCounters["multiCounterName"].Increment("counterName")
	met.MultiCounters["multiCounterName"].Update(&resultsMap)
}
```

