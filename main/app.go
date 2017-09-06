package main

import (
	"github.com/shettyh/threadpool"
	"fmt"
	"time"
)

func main(){
	pool := threadpool.NewThreadPool(2000,100000)
	for i:=0;i<100;i++{
		go RunThread(pool)
	}
	time.Sleep(20*time.Minute)
}

type MyTask struct {
	ID int64
}

func (m *MyTask) Run(){
	fmt.Println("Running my task ",m.ID)
}

func RunThread(pool *threadpool.ThreadPool) {
	for i := 0; i < 500; i++ {
		task := &MyTask{ID: int64(i)}
		pool.Execute(task)
	}
}
