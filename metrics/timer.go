package metrics

import (
	"time"
	"sync/atomic"
)

type Timer struct {
	Name string
	BufferSize uint64 
	ms uint64
	fms uint64
	halt chan bool
	buffer []uint64
}

func (self *Timer) Stop() error {

	self.fms = atomic.LoadUint64(&self.ms)
	self.halt <- true
	self.buffer = append(self.buffer, self.fms)
	if uint64(len(self.buffer)) > self.BufferSize && self.BufferSize > 0 {
		self.buffer = self.buffer[1:]
	}
	return nil
}

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
				time.Sleep(time.Microsecond)
			}
		}
	
	}()

	return nil
}

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


