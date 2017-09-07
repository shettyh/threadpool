package threadpool

// Runnable is interface for the jobs that will be executed by the threadpool
type Runnable interface {
	Run()
}
