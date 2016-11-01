package main

import (
	"sync"
	"time"
)

type Service struct {
	started bool
	stpCh   chan struct{}
	mutex   sync.Mutex
}

func (s *Service) Start() {
	s.stpCh = make(chan struct{})
	go func() {
		s.mutex.Lock()
		s.started = true
		s.mutex.Unlock()
		<-s.stpCh // wait to be closed.
	}()
}
func (s *Service) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.started {
		s.started = false
		close(s.stpCh)
	}
}

func main() {
	s := &Service{}
	s.Start()
	time.Sleep(time.Second) // do some work
	s.Stop()
}
