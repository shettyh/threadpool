package threadpool

// Callable the tasks which returns the output after exit should implement this interface
type Callable interface {
	Call() interface{}
}

// Future is the handle returned after submitting a callable task to the thread pool
type Future struct {
	response chan interface{}
	done     bool
}

// callableTask is internally used to wrap the callable and future together
// So that the worker can send the response back through channel provided in Future object
type callableTask struct {
	Task   Callable
	Handle *Future
}

// Get returns the response of the Callable task when done
// Is is the blocking call it waits for the execution to complete
func (f *Future) Get() interface{} {
	return <-f.response
}

// IsDone returns true if the execution is already done
func (f *Future) IsDone() bool {
	return f.done
}
