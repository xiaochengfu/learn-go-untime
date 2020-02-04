package main

import (
	"log"
	"sync"
	"time"
	"os"
)


//测试sync中的once,使用once.do时，程序只会执行一次,保证并发安全
func main() {
	//未使用sync.once时
	log.Println("--------未使用sync.once时----------")
	for i := range make([]int,10){
		log.Println("循环打印:",i)
	}

	go func() {
		log.Println("并发打印")
	}()

	//使用sync.once时
	var once sync.Once
	log.Println("--------使用sync.once时----------")
	for i := range make([]int,10){
		once.Do(func() {
			log.Println("仅执行一次循环打印:",i)
		})
	}

	go func() {
		once.Do(func() {
			log.Println("这里不会执行到并发打印")
		})
	}()

	//停留一秒，让协程有时间执行
	time.Sleep(1 * time.Second)

	os.Exit(0)
}
