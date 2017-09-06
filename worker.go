package threadpool

type Worker struct {
	jobChannel chan Runnable
	workerPool chan chan Runnable
}

func NewWorker(workerPool chan chan Runnable) *Worker {
	return &Worker{workerPool: workerPool, jobChannel: make(chan Runnable)}
}

func (w Worker) Start() {
	go func() {
		for {
			w.workerPool <- w.jobChannel

			select {
			case job := <-w.jobChannel:
				job.Run()
			}
		}
	}()
}
