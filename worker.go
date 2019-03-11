package threadpool

// Worker type holds the job channel and passed worker threadpool
type Worker struct {
	jobChannel  chan interface{}
	workerPool  chan chan interface{}
	closeHandle chan bool
}

// NewWorker creates the new worker
func NewWorker(workerPool chan chan interface{}, closeHandle chan bool) *Worker {
	return &Worker{workerPool: workerPool, jobChannel: make(chan interface{}), closeHandle: closeHandle}
}

// Start starts the worker by listening to the job channel
func (w Worker) Start() {
	go func() {
		for {
			// Put the worker to the worker threadpool
			w.workerPool <- w.jobChannel

			select {
			// Wait for the job
			case job := <-w.jobChannel:
				// Got the job
				w.executeJob(job)
			case <-w.closeHandle:
				// Exit the go routine when the closeHandle channel is closed
				return
			}
		}
	}()
}

// executeJob executes the job based on the type
func (w Worker) executeJob(job interface{}) {
	// Execute the job based on the task type
	switch task := job.(type) {
	case Runnable:
		task.Run()
		break
	case callableTask:
		response := task.Task.Call()
		task.Handle.done = true
		task.Handle.response <- response
		break
	}
}
