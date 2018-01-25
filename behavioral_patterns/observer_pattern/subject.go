package observer

import (
	"sync"
)

type Subject struct {
	sync.Mutex
	observers []*Observer
	state     int
}

func (s *Subject) GetState() int {
	return s.state
}

func (s *Subject) SetState(state int) {
	s.Lock()
	defer s.Unlock()
	s.state = state
	s.NotifyAllObservers()
}

func (s *Subject) Attach(o *Observer) {
	s.Lock()
	defer s.Unlock()
	s.observers = append(s.observers, o)
}

func (s *Subject) NotifyAllObservers() {
	for _, o := range s.observers {
		o.Update()
	}
}
