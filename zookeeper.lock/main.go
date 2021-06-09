package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	for i := 1; i <= 10; i++ {
		say(i)
	}
}

func say(i int) {
	c, _, err := zk.Connect([]string{"127.0.0.1:12181"}, time.Second) //*10)

	if err != nil {
		panic(err)
	}
	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err = l.Lock()
	if err != nil {
		panic(err)
	}
	println(i, "加锁成功，待处理业务")
	fmt.Println(i, "hi", time.Now())
	// do some thing
	if i < rand.Intn(10) {
		time.Sleep(1 * time.Second)
		_ = l.Unlock()
		println(i, "模拟业务处理成功，并成功释放锁")
	} else {
		time.Sleep(3 * time.Second)
		println(i, "模拟程序异常无法释放锁")
		c.Close()
	}
}
