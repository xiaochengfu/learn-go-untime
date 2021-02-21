package main

import "fmt"

//错误定义
//var(
//	size    :=1024
//	maxSize = size*2
//)

//正确定义
var (
	size    = 1024
	maxSize = size * 2
)

func main() {
	fmt.Println(size, maxSize)
}

//syntax error: unexpected :=, expecting =
