package main

import "fmt"

const a = 100

var b = 200

func main() {
	//fmt.Println(&a,a)
	//上面写法错误&a,常量不同于变量的在运行期分配内存，
	//常量通常会被编译器在预处理阶段直接展开，作为指令数据使用
	fmt.Println(&b, b)
}

//cannot take the address of a
