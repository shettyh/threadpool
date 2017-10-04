package threadpool

import (
	"testing"
	"time"
)

var (
	schedulerpool *ScheduledThreadPool
)

func TestNewScheduledThreadPool(t *testing.T) {
	schedulerpool = NewScheduledThreadPool(2)
}

func TestScheduledThreadPool_Schedule(t *testing.T) {
	task := &TestTask{TestData: &TestData{Val: "pristine"}}
	schedulerpool.ScheduleOnce(task, time.Second*20)

	time.Sleep(5 * time.Second)

	// It should not be changed until 20 secs
	if task.TestData.Val != "pristine" {
		t.Fail()
	}

	time.Sleep(20 * time.Second)

	// It should be changed after 20 secs
	if task.TestData.Val == "pristine" {
		t.Fail()
	}
}

func TestScheduledThreadPool_Close(t *testing.T) {
	schedulerpool.Close()
}
