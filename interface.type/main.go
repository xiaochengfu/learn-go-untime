package main

import (
	"log"
)

type container []interface{}

func (c *container) put(element interface{}) {
	*c = append(*c, element)
}

func (c *container) get() interface{} {
	element := (*c)[0]
	*c = (*c)[1:]
	return element
}

func main() {
	intContainer := &container{}
	intContainer.put(9)
	intContainer.put(2)
	//在把数据取出来时，因为类型是 interface{} ，所以，你还要做一个转型，
	//只有转型成功，才能进行后续操作（因为 interface{}太泛了，泛到什么类型都可以放
	e, ok := intContainer.get().(int) //对已知的类型进行强制转换，interface具有的功能
	//e, ok := intContainer.get().(string) //类型转换
	if !ok {
		log.Print("类型错误")
	} else {
		log.Printf("%v", e)
	}
}
