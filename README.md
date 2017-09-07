# Golang Threadpool implementation
Scalable threadpool implementation using Go to handle the huge network trafic. 

## Usage
- Implement `Runnable` interface for tha task that needs to be executed. For example


  ```
  type MyTask struct { }
   
  func (t *MyTask) Run(){
    // Do your task here
  }
   
  ```
- Create instance of `ThreadPool`
  ```
  pool := threadpool.NewThreadPool(200,1000000)
  ```
- Create Task and execute
  ```
  task:=&MyTask{}
  pool.Execute(task)
  ```
