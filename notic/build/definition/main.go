package main

import "fmt"

func main() {
	type MyInt1 int   //基于一个类型创建一个新类型，称之为definition
	type MyInt2 = int //基于一个类型定义一个别名，成为alias,MyInt2是int类型的别名
	var i int = 9
	//var i1 MyInt1 = i
	//上面一行定义的新类型不能直接赋值
	//报错cannot use i (type int) as type MyInt1 in assignment，需强制转换
	var i1 MyInt1 = MyInt1(i) //强制转换
	var i2 MyInt2 = i         //别名可以直接赋值
	fmt.Println(i1, i2)
}

//打印结果 9 9
