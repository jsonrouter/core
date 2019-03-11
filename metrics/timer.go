package metrics

import (
	"time"
	"sync"
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
	sync.RWMutex
}

// Stop stops the timer
func (self *Timer) Stop() {

	self.Lock()
	self.buffer = append(self.buffer, self.fms)
	self.fms = atomic.LoadUint64(&self.ms)
	self.Unlock()

	self.halt <- true

	self.Lock()
	if uint64(len(self.buffer)) > self.BufferSize && self.BufferSize > 0 {
		self.buffer = self.buffer[1:]
	}
	self.Unlock()
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
func (self *Timer) Update(f func(k string, v interface{})) {

	self.RLock()
	bufferSize := uint64(len(self.buffer))
	r := sum(self.buffer...)
	self.RUnlock()

	if (bufferSize != 0) {
		r = r / bufferSize
	}

	f(self.Name, r)
}
