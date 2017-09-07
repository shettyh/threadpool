package threadpool

//ThreadPool type for holding the workers and handle the job requests
type ThreadPool struct {
	QueueSize   int64
	NoOfWorkers int

	jobQueue   chan Runnable
	workerPool chan chan Runnable
}

// NewThreadPool creates threadpool
func NewThreadPool(noOfWorkers int, queueSize int64) *ThreadPool {
	threadPool := &ThreadPool{QueueSize: queueSize, NoOfWorkers: noOfWorkers}
	threadPool.jobQueue = make(chan Runnable, queueSize)
	threadPool.workerPool = make(chan chan Runnable, noOfWorkers)

	threadPool.createPool()
	return threadPool
}

// Execute submits the job to available worker
func (t *ThreadPool) Execute(task Runnable) {
	t.jobQueue <- task
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
			go func(job Runnable) {
				jobChannel := <-t.workerPool
				jobChannel <- job
			}(job)
		}
	}
}
