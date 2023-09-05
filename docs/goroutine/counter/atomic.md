# Atomic counter

```go
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

```

<img height="360" src="https://media.giphy.com/media/O6nT9DSoiUVYQ/giphy.gif" width="389" alt="?"/>