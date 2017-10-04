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
	threadpool *ThreadPool
)

func TestNewThreadPool(t *testing.T) {
	threadpool = NewThreadPool(NumberOfWorkers, QueueSize)
}

func TestThreadPool_Execute(t *testing.T) {
	data := &TestData{Val: "pristine"}
	task := &TestTask{TestData: data}
	threadpool.Execute(task)

	time.Sleep(2 * time.Second)
	fmt.Println("")

	if data.Val != "changed" {
		t.Fail()
	}
}

func TestThreadPool_ExecuteFuture(t *testing.T) {
	task := &TestTaskFuture{}
	handle := threadpool.ExecuteFuture(task)
	response := handle.Get()
	if !handle.IsDone() {
		t.Fail()
	}
	fmt.Println("Thread done ", response)
}

func TestThreadPool_Close(t *testing.T) {
	threadpool.Close()
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

type TestTaskFuture struct{}

func (t *TestTaskFuture) Call() interface{} {
	return "Done"
}
