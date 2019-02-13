package metrics

import (
	"sync"
	"fmt"
)

type Counter struct {
	Name string
	t uint64
	sync.RWMutex
}

func (self *Counter) GetValue() uint64 {
	self.RLock()
	defer self.RUnlock()
	return self.t
}

func (self *Counter) Reset() {
	self.Lock()
	defer self.Unlock()
	self.t = 0
}

func (self *Counter) Increment() {
	self.Lock()
	defer self.Unlock()
	self.t += 1
}

func (self *Counter) Update(results *map[string]interface{}) error {
fmt.Println("Update c")
	self.Lock()
	defer self.Unlock()
	res := *results
	res[self.Name] = self.t
	return nil
}
