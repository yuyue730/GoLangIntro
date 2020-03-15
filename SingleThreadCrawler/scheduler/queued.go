package scheduler

import "GoLangIntro/SingleThreadCrawler/engine"

// Queued Scheduler has a channel of engine and a channel of channel of request
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {

}

// When QueuedScheduler is run, it will run into two scenarios
// 1. When a new worker or request comes it, it will add that worker or request
// item at the back of the queue.
// 2. When we need the worker to work on the request, we pop both front item from
// Request and Worker Queue and feed request item into worker item which is a
// Channel of request, in `engine/worker`, the worker function is going to fecth
// and parse the request.
func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
