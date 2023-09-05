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
