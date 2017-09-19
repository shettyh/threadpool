package threadpool

type Callable interface {
	Call() interface{}
}

type CallableTask struct {
	Task Callable
	Handle *Future
}

type Future struct {
	response chan interface{}
	done bool
}

//Blocking call
func (f *Future) Get() interface{}{
	return <-f.response
}

func (f *Future) IsDone() bool{
	return f.done
}
