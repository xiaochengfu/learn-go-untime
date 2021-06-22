package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup


func worker(exitChan chan struct{}) {
LOOP:
	for  {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}
	wg.Done()
}


func main() {
	exitChan := make(chan struct{})
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
	exitChan <- struct{}{}      // 给子goroutine发送退出信号
	close(exitChan)
	wg.Wait()
	fmt.Println("over")
}

