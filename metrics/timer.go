package metrics

import (
	"time"
	"sync/atomic"
)

// Timer is a simple timer used to measure time taken to perform given tasks in microseconds
type Timer struct {
	Name string
	BufferSize uint64 
	ms uint64
	fms uint64
	halt chan bool
	buffer []uint64
}

// Stop stops the timer
func (self *Timer) Stop() error {

	self.fms = atomic.LoadUint64(&self.ms)
	self.halt <- true
	self.buffer = append(self.buffer, self.fms)
	if uint64(len(self.buffer)) > self.BufferSize && self.BufferSize > 0 {
		self.buffer = self.buffer[1:]
	}
	return nil
}

// Start starts the timer. Place what needs to be timed in between Start() and Stop()
func (self *Timer) Start() error {
	
	self.ms = 0
	self.halt = make(chan bool)

	go func() {
		for {	
			select {
			case <- self.halt:
				return
			default:
				atomic.AddUint64(&self.ms, 1)
				time.Sleep(time.Nanosecond)
			}
		}
	
	}()

	return nil
}

// Update is called to save the values of the timer into the results map. 
// You can pass any map[string]Interface{} to store results including the provide Results
// map on the main Metrics struct
func (self *Timer) Update(results *map[string]interface{}) error {
	res := *results
	bufferSize := uint64(len(self.buffer))
	r := sum(self.buffer...)

	if (bufferSize != 0) {
		r = r / bufferSize
	}

	res[self.Name] = r

	return nil
}


