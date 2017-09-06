package main

import (
	"github.com/shettyh/threadpool"
	"fmt"
	"time"
)

func main(){
	pool := threadpool.NewThreadPool(20,2000)

	for i:=0;i<5;i++ {
		task:=&MyTask{ID:int64(i)}
		pool.Execute(task)
	}

	time.Sleep(20*time.Minute)
}

type MyTask struct {
	ID int64
}

func (m *MyTask) Run(){
	fmt.Println("Running my task ",m.ID)
}
