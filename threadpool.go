package threadpool

//ThreadPool type for holding the workers and handle the job requests
type ThreadPool struct {
	QueueSize   int64
	NoOfWorkers int

	jobQueue   chan interface{}
	workerPool chan chan interface{}
}

// NewThreadPool creates thread pool
func NewThreadPool(noOfWorkers int, queueSize int64) *ThreadPool {
	threadPool := &ThreadPool{QueueSize: queueSize, NoOfWorkers: noOfWorkers}
	threadPool.jobQueue = make(chan interface{}, queueSize)
	threadPool.workerPool = make(chan chan interface{}, noOfWorkers)

	threadPool.createPool()
	return threadPool
}

// Execute submits the job to available worker
func (t *ThreadPool) Execute(task Runnable) {
	// Add the task to the job queue
	t.jobQueue <- task
}

// ExecuteFuture will submit the task to the threadpool and return the response handle
func (t *ThreadPool) ExecuteFuture(task Callable) chan *Future {
	futureTask:= CallableTask{Task:task,Handle:make(chan *Future)}
	t.jobQueue <- futureTask
	return futureTask.Handle
}

// createPool creates the workers and start listening on the jobQueue
func (t *ThreadPool) createPool() {
	for i := 0; i < t.NoOfWorkers; i++ {
		worker := NewWorker(t.workerPool)
		worker.Start()
	}

	go t.dispatch()

}

// dispatch listens to the jobqueue and handles the jobs to the workers
func (t *ThreadPool) dispatch() {
	for {
		select {
		case job := <-t.jobQueue:
			// Got job
			go func(job interface{}) {
				//Find a worker for the job
				jobChannel := <-t.workerPool
				//Submit job to the worker
				jobChannel <- job
			}(job)
		}
	}
}
