package engine

import "log"

// Struct that defines how concurrent engine works. It takes in a struct under
// scheduler interface and number of worker needed.
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemsCount  int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// in := make(chan Request) Only need for Simple Scheduler, the in channel
	// will be declared in `createWorker`
	out := make(chan ParseResult)

	// Pushing the in channel as the master work channel
	// e.Scheduler.ConfigureMasterWorkerChan(in) Only need for Simple Scheduler
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		// createWorker(in, out) Only need for Simple Scheduler
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		// Listen to what createWorker function feed into out channel and keep
		// track of these ParseResults
		result := <-out
		for _, item := range result.Items {
			e.ItemsCount++
			log.Printf("Got Items: No. %d, Content %v", e.ItemsCount, item)
		}

		if e.ItemsCount >= 1000 {
			break
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
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

// Declare in channel Request and feed it to QueuedScheduler struct, in the
// goroutine, keeps reading from in channel and letring worker perform on the
// in channel request, feed the result to the out channel
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
