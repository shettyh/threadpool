package threadpool

type Callable interface {
	Call() interface{}
}

type CallableTask struct {
	Task Callable
	Handle chan *Future
}

type Future struct {
	response interface{}
}

func (f *Future) Get() interface{}{
	return f.response
}
