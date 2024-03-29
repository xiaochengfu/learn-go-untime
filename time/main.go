package main

import (
	"fmt"
	"time"
)

func testTimer1() {
	go func() {
		fmt.Println("test timer1")
	}()
}

func ticker() {
	timer1 := time.NewTicker(2 * time.Second)
	defer timer1.Stop()
	for {
		select {
		case <-timer1.C:
			testTimer1()
		}
	}
}

func timer() {
	time2 := time.NewTimer(5 * time.Second)
	defer time2.Stop()
	select {
	case <-time2.C:
		fmt.Println("5s了")
	}
}

func main() {
	startTime := time.Now()
	fmt.Println(time.Now()) //2021-02-15 10:54:15.598193 +0800 CST m=+0.000074813
	//时间格式化
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) //2021-02-15 10:54:15

	//秒时间戳 10位
	fmt.Println(time.Now().Unix()) //1613357655
	//指定时间转时间戳
	pTime, _ := time.Parse("2006-01-02 15:04", "2021-12-24 10:00")
	fmt.Println(pTime.Unix()) //1640340000
	//毫秒时间戳 13位
	fmt.Println(time.Now().UnixNano() / 1e6) //1613357655598
	//纳秒时间戳 19位
	fmt.Println(time.Now().UnixNano()) //1613357655598425000

	//时间戳转时间格式
	fmt.Println(time.Unix(1613355650, 0).Format("2006-01-02 15:04:05")) //2021-02-15 10:20:50

	//日期格式转时间戳
	pDate, _ := time.Parse("2006-01-02", "2021-12-24")
	fmt.Println(pDate.Unix()) //1640304000

	endTime := time.Now()
	fmt.Println("执行耗时：", endTime.Sub(startTime))

	go ticker() //每2秒执行一次
	go timer()  //5秒后执行
	time.Sleep(10 * time.Second)
}
