package worker

import "fmt"

var queue = make(chan Request, 10)

// start starts the worker -- ALWAYS call this as a goroutine
func start() {
	defer close(queue)

	// infinite loop
	for {
		// wait for message in worker queue
		msg := <-queue

		fmt.Println("Processing message:", msg.Id)

		// offset message handling
		res, err, shutdown := handeMessage(msg)

		// on error, return standard error message and continue
		if err != nil {
			msg.ResCh <- *workerError(msg.Id, err)
			continue
		}

		// send response to channel
		msg.ResCh <- res

		// if shutdown is true, break out of loop
		if shutdown {
			fmt.Println("Worker shutting down")
			return
		}
	}
}

// New creates a new worker and returns the channel to send messages to
// the channel is unique per worker and if messages processing is sequential
func New() chan Request {
	go start()
	return queue
}
