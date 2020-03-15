package engine

import "log"

// Struct that defines how concurrent engine works. It takes in a struct under
// scheduler interface and number of worker needed.
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// in := make(chan Request) Only need for Simple Scheduler
	out := make(chan ParseResult)

	// Pushing the in channel as the master work channel
	// e.Scheduler.ConfigureMasterWorkerChan(in) Only need for Simple Scheduler
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		// createWorker(in, out) Only need for Simple Scheduler
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got Items: %v", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// Read `in` channel into `request`, call `worker` on it and push result into
// the `out` channel Only need for Simple Scheduler
// func createWorker(in chan Request, out chan ParseResult) {
// 	go func() {
// 		for {
// 			request := <-in
// 			result, err := worker(request)
// 			if err != nil {
// 				continue
// 			}
// 			out <- result
// 		}
// 	}()
// }

func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
