package main

import "time"

type Service struct {
	started bool
	stpCh   chan struct{}
}

func (s *Service) Start() {
	s.stpCh = make(chan struct{})
	go func() {
		s.started = true
		<-s.stpCh // wait to be closed.
	}()
}
func (s *Service) Stop() {
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
