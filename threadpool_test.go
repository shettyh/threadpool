package threadpool

import (
	"fmt"
	"testing"
	"time"
)

const (
	NumberOfWorkers = 20
	QueueSize       = int64(1000)
)

var (
	Pool *ThreadPool
)

func TestNewThreadPool(t *testing.T) {
	Pool = NewThreadPool(NumberOfWorkers, QueueSize)
}

func TestThreadPool_Execute(t *testing.T) {
	data := &TestData{Val: "pristine"}
	task := &TestTask{TestData: data}
	Pool.Execute(task)

	time.Sleep(2 * time.Second)
	fmt.Println("")

	if data.Val != "changed" {
		t.Fail()
	}
}

func TestThreadPool_ExecuteFuture(t *testing.T) {
	task:= &TestTaskFuture{}
	handle:=Pool.ExecuteFuture(task)
	response := handle.Get()
	fmt.Println("Thread done ",response)
}

type TestTask struct {
	TestData *TestData
}

type TestData struct {
	Val string
}

func (t *TestTask) Run() {
	fmt.Println("Running the task")
	t.TestData.Val = "changed"
}

type TestTaskFuture struct {}

func (t *TestTaskFuture) Call() interface{} {
	return "Done"
}
