package counter

import (
	"sync/atomic"
)

func NewAtomicCounter() Counter {
	return &atomicCounter{}
}

type atomicCounter struct {
	counter atomic.Int64
}

func (s *atomicCounter) Inc() int {
	return int(s.counter.Add(1))
}

func (s *atomicCounter) Count() int {
	return int(s.counter.Load())
}
