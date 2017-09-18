package threadpool

import (
	"testing"
	"time"
)

var (
	pool *ScheduledThreadPool
)

func TestNewScheduledThreadPool(t *testing.T) {
	pool = NewScheduledThreadPool(2)
}

func TestScheduledThreadPool_Schedule(t *testing.T) {
	task := &TestTask{TestData:&TestData{Val:"pristine"}}
	pool.Schedule(task,time.Second*20)

	time.Sleep(5*time.Second)

	if task.TestData.Val != "pristine" {
		t.Fail()
	}

	time.Sleep(20*time.Second)

	if task.TestData.Val == "pristine" {
		t.Fail()
	}
}
