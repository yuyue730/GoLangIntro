package scheduler

import "GoLangIntro/SingleThreadCrawler/engine"

// Simple Scheduler, with a channel of engine.Request
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// Pass engine.Request into the workChan member in a coroutine.
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {}
