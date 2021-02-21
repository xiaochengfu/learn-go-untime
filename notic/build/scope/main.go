package main

import (
	"fmt"
)

func Repeat() (i int) {
	for i := 0; i <= 5; i++ {
		//for循环内存的代码块，对i重写赋值了，作用域只在其代码块中生效
		fmt.Println("循环内i=", i)
		//循环内i= 0
		//循环内i= 1
		//循环内i= 2
		//循环内i= 3
		//循环内i= 4
		//循环内i= 5
	}
	return i //返回默认值0
}

func DoIf(m bool) (j int) {
	if m {
		//if循环内存的代码块，对i重写赋值了，作用域只在其代码块中生效
		//跟外层的返回值j没有关系了
		j := 100
		j++
	} else {
		j = 100
		j++
	}
	return j
}

func main() {
	fmt.Println("-------for---------")
	//如果变量是在一个代码块，比如 for / if 中，那么这个变量的的作用域就在该代码块
	// 如果变量是在一个for代码块
	result := Repeat()
	fmt.Println(result)
	//结果
	//循环内i= 0
	//循环内i= 1
	//循环内i= 2
	//循环内i= 3
	//循环内i= 4
	//循环内i= 5
	//0

	fmt.Println("--------if--------")
	// 如果变量是在一个for代码块
	result2 := DoIf(true)
	result3 := DoIf(false)
	fmt.Println(result2) //0
	fmt.Println(result3) //101
}
