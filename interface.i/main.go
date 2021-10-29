package main

import "fmt"

type Study interface {
	Listen(message string) string
	i() //定义一个小写方法，来约束此接口只能在本包内使用
}

//接口类型检查,study需要实现Study接口
var _ Study = (*study)(nil)

type study struct {
	Name string
}

func (s study) i() {
	panic("implement me")
}

func (s study) Listen(message string) string {
	panic("implement me")
}

func main() {
	fmt.Println("hello world")
}
