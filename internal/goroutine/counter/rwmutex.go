package counter

import (
	"sync"
)

func NewRWMutexCounter() Counter {
	return &rwmutexCounter{}
}

type rwmutexCounter struct {
	counter int

	lock sync.RWMutex
}

func (s *rwmutexCounter) Inc() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.counter++
	return s.counter
}

func (s *rwmutexCounter) Count() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.counter
}
