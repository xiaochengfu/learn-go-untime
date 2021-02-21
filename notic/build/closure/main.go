package main

import "fmt"

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		i := i //i重写赋值没有搞清楚
		funs = append(funs, func() {
			//匿名函数的优越性在于可以直接使用函数内的变量，不必申明
			println(&i, i)
		})
	}
	return funs
}

func main() {
	funs := test()
	fmt.Println(funs)
	for _, f := range funs {
		f()
	}
}

//0xc000018090 0
//0xc000018098 1
