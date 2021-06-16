package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)


var syncM = sync.Map{}

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

// map 非并发安全测试
func mapTest() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 并发安全测试
func syncMapTest() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			syncM.Store(key, n)
			value, _ := syncM.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

	syncM.LoadOrStore(100,"hp")
	syncM.Delete(100)

	syncM.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:",key,value)
		return true
	})
}

func main() {
	//mapTest()
	syncMapTest()
}

