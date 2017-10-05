# Golang Threadpool implementation
[![Build Status](https://travis-ci.org/shettyh/threadpool.svg?branch=master)](https://travis-ci.org/shettyh/threadpool)
[![codecov](https://codecov.io/gh/shettyh/threadpool/branch/master/graph/badge.svg)](https://codecov.io/gh/shettyh/threadpool)

Scalable threadpool implementation using Go to handle the huge network trafic. 

## Install

`go get github.com/shettyh/threadpool`

## Usage

### Threadpool
- Implement `Runnable` interface for tha task that needs to be executed. For example


  ```
  type MyTask struct { }
   
  func (t *MyTask) Run(){
    // Do your task here
  }
   
  ```
- Create instance of `ThreadPool` with number of workers required and the task queue size
  ```
  pool := threadpool.NewThreadPool(200,1000000)
  ```
- Create Task and execute
  ```
  task:=&MyTask{}
  pool.Execute(task)
  ```
- Using `Callable` task
  ```
  type MyTaskCallable struct { }
  
  func (c *MyTaskCallable) Call() interface{} {
    //Do task 
    return result
  }
  
  //Execute callable task
  task := &MyTaskCallable{}
  future := pool.ExecuteFuture(task)
  
  //Check if the task is done
  isDone := future.IsDone() // true/false
  
  //Get response , blocking call
  result := future.Get()
  
  ```
- Close the pool
  ```
  pool.Close()
  ```

### Scheduled threadpool

- Create instance of `ScheduledThreadPool` with number of workers required
  ```
  schedulerPool:= threadpool.NewScheduledThreadPool(10)
  ```
- Create Task and schedule
  ```
  task:=&MyTask{}
  pool.ScheduleOnce(task, time.Second*20) // Time delay is in seconds only as of now
  ```
- Close the pool
  ```
  pool.Close()
  ```
