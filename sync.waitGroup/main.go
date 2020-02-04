package main

import (
	"log"
	"sync"
	"time"
)

//测试sync.waitGroup
//等待所有任务执行完成后，才去执行下一步操作
//例如：从1加到100等于多少？结果等于5050，算出结果后，我再乘以10，最后等于50500
//这里我起9个携程，分别计算，计算完后，再乘以10
func main() {
	log.Println("------------使用sync.waitGroup------------")
	startTime := time.Now().Unix()
	endNum := 100
	var wg sync.WaitGroup
	var total int
	for i := 1; i <= endNum/10; i++ {
		start := (i-1)*10+1
		wg.Add(1)
		go func(i int) {
			sum := calculate(start,i*10)
			//log.Println(start,"->",i*10,"=",sum)
			total += sum
			wg.Done()
		}(i)
	}
	wg.Wait()
	endTime := time.Now().Unix()
	log.Println("从1加到",endNum,"在乘以10，结果等于",total*10,"用时",endTime-startTime,"s")

	log.Println("------------不使用sync.waitGroup------------")
	startTime2 := time.Now().Unix()
	var total2 int
	for i := 1; i <= endNum/10; i++ {
		start := (i-1)*10+1
		go func(i int) {
			sum := calculate(start,i*10)
			//log.Println(start,"->",i*10,"=",sum)
			total2 += sum
		}(i)
	}
	//需要预留时间，让协程有机会执行,如果不用channel,很难确定应该预留几秒
	time.Sleep(1*time.Second)
	endTime2 := time.Now().Unix()
	log.Println("从1加到",endNum,"在乘以10，结果等于",total2*10,"用时",endTime2-startTime2,"s")


}

//计算从几加到几的结果
func calculate(start int, end int) (sum int){
	for i:=start;i<=end;i++{
		sum += i
	}
	return sum
}
