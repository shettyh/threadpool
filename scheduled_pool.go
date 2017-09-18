package threadpool

import (
	"sync"
	"math"
	"nokia.com/nas/services/telemetry_service_go/ds"
)

type ScheduledThreadPool struct {
	workers chan chan Runnable
	jobQueue chan Runnable
	tasks *sync.Map
	noOfWorkers int
	counter uint64
	queueSize int64
	counterLock sync.Mutex
}

func NewScheduledThreadPool(noOfWorkers int) *ScheduledThreadPool {
	pool := &ScheduledThreadPool{}
	pool.noOfWorkers=noOfWorkers
	pool.queueSize = math.MaxInt32
	pool.workers = make(chan chan Runnable,noOfWorkers)
	pool.jobQueue = make(chan Runnable,pool.queueSize)
	pool.tasks = new(sync.Map)
	pool.createPool()
	return pool
}

func (stf *ScheduledThreadPool) createPool(){
	for i:=0;i<stf.noOfWorkers; i++ {
		worker:= NewWorker(stf.workers)
		worker.Start()
	}

	go stf.dispatch()
}

func (stf *ScheduledThreadPool) dispatch(){
	go stf.intervalRunner()
}

func (stf *ScheduledThreadPool) intervalRunner(){
	stf.updateCounter()
	currentTasksToRun,ok:=stf.tasks.Load(stf.counter)

	if ok {
		currentTasksSet := currentTasksToRun.(*ds.Set)

		for _,val:= range currentTasksSet.GetAll() {
			task:=val.(Runnable)
			worker:=<-stf.workers
			worker <- task
		}
	}
}

func (stf *ScheduledThreadPool) updateCounter(){
	stf.counterLock.Lock()
	stf.counter++
	stf.counterLock.Unlock()
}



