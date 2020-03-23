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
	handle, _ := threadpool.ExecuteFuture(task)
	response := handle.Get()
	if !handle.IsDone() {
		t.Fail()
	}
	fmt.Println("Thread done ", response)
}

func TestThreadPool_Close(t *testing.T) {
	threadpool.Close()
}


func TestQueueFullError(t *testing.T) {
	threadpool := NewThreadPool(0, 1)

	data := &TestData{Val: "pristine"}
	task := &TestTask{TestData: data}

	err := threadpool.Execute(task)
	if err != nil {
		t.Fail()
	}

	err = threadpool.Execute(task)
	if err == nil || err != ErrQueueFull {
		t.Fail()
	}

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

type TestLongTask struct { }

func (t TestLongTask) Run() {
	time.Sleep(5 * time.Second)
}

type TestTaskFuture struct{}

func (t *TestTaskFuture) Call() interface{} {
	return "Done"
}
