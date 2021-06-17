package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	lock.Lock()   // 加互斥锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
	lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	lock.Lock()                  // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	lock.Unlock()                // 解互斥锁
	wg.Done()
}

func writeRW() {
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	wg.Done()
}

func readRW() {
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	wg.Done()
}

func mutexLock() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func RwLock() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeRW()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go readRW()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func main() {
	mutexLock()
	RwLock()
}

//读写锁适合读多写少的场景，性能要高于互斥锁
//互斥锁执行时间：1.345988348s
//读写锁执行时间：111.678589ms