package main

import (
	"fmt"
	"math/rand"
)

//写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，
//另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

func main() {
	//方式一
	ch1 := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			select {
			case ch1 <- rand.Intn(100):
			case <-done:
				return
			default:
			}
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("rand num = ", <-ch1)
		}
		done <- true
		return
	}()
	<-done

	fmt.Println("-----------------")
	//方式二
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(100)
		}

		close(ch)
	}()

	for {
		data, isClose := <-ch
		if !isClose {
			return
		}

		fmt.Println(data)
	}

}
