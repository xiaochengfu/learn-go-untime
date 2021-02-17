package main

import "log"

func echo(nums []int) <-chan int {
	//相关数组还没有定义时，可以用make来创建一个切片
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//平方函数
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

//过滤奇数函数
func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 != 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

//求和函数
func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

type EchoFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

//已代理的形式实现类似管道式的方法
func pipeline(nums []int, echo EchoFunc, pipeFns ...PipeFunc) <-chan int {
	ch := echo(nums)
	for _, fn := range pipeFns {
		ch = fn(ch)
	}
	//或者换种写法
	//for i := range pipeFns {
	//	ch = pipeFns[i](ch)
	//}
	return ch
}

func main() {
	//一个由数字1到10组成的切片,完整写法为 var nums = []int{1,2,3,4,5,6,7,8,9,10}[:]，[:]代表整个数组
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//range可以操作chan
	for n := range sum(sq(odd(echo(nums)))) {
		log.Println(n)
	}

	//上面的代码类似于我们执行了 Unix/Linux 命令： echo $nums | sq | sum。
	//同样，如果你不想有那么多的函数嵌套，就可以使用一个代理函数来完成。
	for m := range pipeline(nums, echo, odd, sq, sum) {
		log.Println(m)
	}
}
