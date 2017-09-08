package threadpool

// Worker type holds the job channel and passed worker pool
type Worker struct {
	jobChannel chan Runnable
	workerPool chan chan Runnable
}

// NewWorker creates the new worker
func NewWorker(workerPool chan chan Runnable) *Worker {
	return &Worker{workerPool: workerPool, jobChannel: make(chan Runnable)}
}

// Start starts the worker by listening to the job channel
func (w Worker) Start() {
	go func() {
		for {
			// Put the worker to the worker pool
			w.workerPool <- w.jobChannel

			select {
			// Wait for the job
			case job := <-w.jobChannel:
				// Execute the job
				job.Run()
			}
		}
	}()
}
