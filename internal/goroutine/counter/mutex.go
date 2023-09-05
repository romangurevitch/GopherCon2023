package counter

import (
	"sync"
)

func NewMutexCounter() Counter {
	return &mutexCounter{}
}

type mutexCounter struct {
	counter int

	lock sync.Mutex
}

func (s *mutexCounter) Inc() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.counter++
	return s.counter
}

func (s *mutexCounter) Count() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.counter
}
