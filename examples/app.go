package examples

import (
	"fmt"
	"github.com/shettyh/threadpool"
	"time"
)

func main() {
	pool := threadpool.NewThreadPool(2000, 100000)
	time.Sleep(20 * time.Minute)
	task := &myTask{ID: 123}
	pool.Execute(task)
}

type myTask struct {
	ID int64
}

func (m *myTask) Run() {
	fmt.Println("Running my task ", m.ID)
}
