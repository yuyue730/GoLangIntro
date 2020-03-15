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

// Configure the first and the master worker channel
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
