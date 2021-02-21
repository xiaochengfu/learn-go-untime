package main

import "fmt"

var j int

func comeOn() {
	fmt.Println("我跳一遍")
	j++
	if j > 5 {
		panic("不要来回跳了")
	}
}

func main() {
	//for i:=0;i<10 ;i++  {
	//loop:
	//	println(i)
	//}
	//goto loop

	//以上写法会报错 goto loop jumps into block starting at
	//goto不能跳转到其他函数或者内层代码

	for i := 0; i < 10; i++ {
		println(i)
	}
loop:
	comeOn() //来回goto相当于循环了
	goto loop
}
