package main

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Stringable interface {
	ToString() string
}

func (c Country) ToString() string {
	return "Country = " + c.Name
}

func (c City) ToString() string {
	return "City = " + c.Name
}

func printStr(p Stringable) {
	fmt.Println(p.ToString())
}

func main() {
	d1 := Country{"USA"}
	d2 := City{"Los Angeles"}
	//接口接受任意类型的传参
	printStr(d1)
	printStr(d2)

	//在这段代码中，我们可以看到，我们使用了一个叫Stringable 的接口，
	//我们用这个接口把“业务类型” Country 和 City 和“控制逻辑” Print() 给解耦了。
	//于是，只要实现了Stringable 接口，都可以传给 PrintStr() 来使用。
}
