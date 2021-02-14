package main

import "fmt"

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s Square) Area() int {
	panic("implement me")
}

func (s* Square) Sides() int {
	return 4
}

func main() {

	//Square 并没有实现 Shape 接口的所有方法，程序虽然可以跑通，
	//但是这样的编程方式并不严谨，如果我们需要强制实现接口的所有方法

	//声明一个 _ 变量（没人用）会把一个 nil 的空指针从 Square 转成 Shape，
	//这样，如果没有实现完相关的接口方法，编译器就会报错
	var _ Shape = (*Square)(nil) //声明的变量_是一个接口类型Shape, 而square如果没有实现接口全部的方法，会报错

	s := Square{len: 5}
	fmt.Printf("%d\n",s.Sides())
}

