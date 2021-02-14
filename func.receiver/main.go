package main

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

//成员函数，假设结构体是类的一种简化，那类的方法则就是像下面这样的成员函数，
//receiver接收者c,是Country结构体的变量，而PrintStr则是作用在接收体上的函数
func (c Country) PrintStr() {
	fmt.Println(c.Name)
}

//接收者的类型可以是指针
func (c *Country) PrintStr2() {
	c.Name = "Zhong-Guo"
	fmt.Println(c.Name)
}

func (c City) PrintStr()  {
	fmt.Println(c.Name)
}

func main() {
	c1 := Country{"China"}
	c1.PrintStr2() //Zhong-Guo

	c2 := City{"Beijing"}
	c2.PrintStr() //Beijing

	c1.PrintStr() //Zhong-Guo
}
