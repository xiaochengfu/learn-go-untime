package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			//输出的结果取决于外层循环到哪里了，所有每次打印的结果都是不一样的
			fmt.Println(i)
			wg.Done()
		}()
	}
	//等待所有的任务都执行完才结束
	wg.Wait()
	fmt.Println("-----------------------")
	runtime.GOMAXPROCS(1) //设置执行时的最大cpu数，设为1则每次打印的结果是一致的
	var wg2 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg2.Done()
		}(i)
	}
	wg2.Wait()
}

//执行结果：
//10
//10
//10
//3
//10
//10
//10
//10
//10
//10
//-----------------------
//9
//0
//1
//2
//3
//4
//5
//6
//7
//8
