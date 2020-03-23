package threadpool

import "fmt"

var (
	ErrQueueFull = fmt.Errorf("queue is full, not able add the task")
)

//ThreadPool type for holding the workers and handle the job requests
type ThreadPool struct {
	QueueSize   int64
	NoOfWorkers int

	jobQueue    chan interface{}
	workerPool  chan chan interface{}
	closeHandle chan bool // Channel used to stop all the workers
}

// NewThreadPool creates thread threadpool
func NewThreadPool(noOfWorkers int, queueSize int64) *ThreadPool {
	threadPool := &ThreadPool{QueueSize: queueSize, NoOfWorkers: noOfWorkers}
	threadPool.jobQueue = make(chan interface{}, queueSize)
	threadPool.workerPool = make(chan chan interface{}, noOfWorkers)
	threadPool.closeHandle = make(chan bool)
	threadPool.createPool()
	return threadPool
}

func (t *ThreadPool) submitTask( task interface{}) error {
	// Add the task to the job queue
	if len(t.jobQueue) == int(t.QueueSize) {
		return ErrQueueFull
	}
	t.jobQueue <- task
	return nil
}

// Execute submits the job to available worker
func (t *ThreadPool) Execute(task Runnable) error {
	return t.submitTask(task)
}

// ExecuteFuture will submit the task to the threadpool and return the response handle
func (t *ThreadPool) ExecuteFuture(task Callable) (*Future, error) {
	// Create future and task
	handle := &Future{response: make(chan interface{})}
	futureTask := callableTask{Task: task, Handle: handle}
	err := t.submitTask(futureTask)
	if err != nil {
		return nil, err
	}
	return futureTask.Handle, nil
}

// Close will close the threadpool
// It sends the stop signal to all the worker that are running
//TODO: need to check the existing /running task before closing the threadpool
func (t *ThreadPool) Close() {
	close(t.closeHandle) // Stops all the routines
	close(t.workerPool)  // Closes the Job threadpool
	close(t.jobQueue)    // Closes the job Queue
}

// createPool creates the workers and start listening on the jobQueue
func (t *ThreadPool) createPool() {
	for i := 0; i < t.NoOfWorkers; i++ {
		worker := NewWorker(t.workerPool, t.closeHandle)
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
			func(job interface{}) {
				//Find a worker for the job
				jobChannel := <-t.workerPool
				//Submit job to the worker
				jobChannel <- job
			}(job)

		case <-t.closeHandle:
			// Close thread threadpool
			return
		}
	}
}
