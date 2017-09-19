package threadpool

// Worker type holds the job channel and passed worker pool
type Worker struct {
	jobChannel chan interface{}
	workerPool chan chan interface{}
}

// NewWorker creates the new worker
func NewWorker(workerPool chan chan interface{}) *Worker {
	return &Worker{workerPool: workerPool, jobChannel: make(chan interface{})}
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
				// Execute the job based on the task type
				switch job.(type) {
				case Runnable:
					job.(Runnable).Run()
					break
				case callableTask:
					task := job.(callableTask)
					response := task.Task.Call()
					task.Handle.done=true
					task.Handle.response <- response
					break
				}
			}
		}
	}()
}
