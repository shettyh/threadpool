package threadpool

type ThreadPool struct {
	QueueSize   int64
	NoOfWorkers int

	jobQueue   chan Runnable
	workerPool chan chan Runnable
}

func NewThreadPool(noOfWorkers int, queueSize int64) *ThreadPool {
	threadPool := &ThreadPool{QueueSize: queueSize, NoOfWorkers: noOfWorkers}
	threadPool.jobQueue = make(chan Runnable, queueSize)
	threadPool.workerPool = make(chan chan Runnable, noOfWorkers)

	threadPool.createPool()
	return threadPool
}

func (t *ThreadPool) Execute(task Runnable) {
	t.jobQueue <- task
}

func (t *ThreadPool) createPool() {
	for i := 0; i < t.NoOfWorkers; i++ {
		worker := NewWorker(t.workerPool)
		worker.Start()
	}

	go t.dispatch()

}

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
