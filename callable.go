package threadpool

type Callable interface {
	Call() interface{}
}

type Future struct {
	response chan interface{}
	done bool
}

type callableTask struct {
	Task Callable
	Handle *Future
}

//Blocking call
func (f *Future) Get() interface{}{
	return <-f.response
}

func (f *Future) IsDone() bool{
	return f.done
}
