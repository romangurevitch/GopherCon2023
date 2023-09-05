# Basic counter

```go
package counter

func NewBasicCounter() Counter {
	return &basicCounter{}
}

type basicCounter struct {
	counter int
}

func (s *basicCounter) Inc() int {
	s.counter++
	return s.counter
}

func (s *basicCounter) Count() int {
	return s.counter
}
```

<img height="360" src="https://media.giphy.com/media/APqEbxBsVlkWSuFpth/giphy.gif" width="389" alt="?"/>