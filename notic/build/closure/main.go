package main

import "fmt"

func loop() {
	var addressI *int
	for i := 0; i < 2; i++ {
		fmt.Println(i, &i)
		addressI = &i

	}
	fmt.Println(addressI, *addressI)
}

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		fmt.Println("before:", &i, i)
		//i := i //这里要重新赋值，修改掉i的内存地址，f()每次调用才是不同的值
		fmt.Println("end:", &i, i)
		funs = append(funs, func() {
			//匿名函数的优越性在于可以直接使用函数内的变量，不必申明
			println(&i, i)
		})
	}
	return funs
}

func main() {
	loop()
	fmt.Println("-----------")
	funs := test()
	for _, f := range funs {
		f()
	}
}
