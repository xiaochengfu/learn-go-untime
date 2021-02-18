package main

import (
	"fmt"
	"log"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func (u *User) SetName(name string) {
	u.Name = name
	fmt.Printf("设置用户名称为:%v\n", name)
}

func GetFieldAndMethod(input interface{}) {
	//获取类型名称
	userType := reflect.TypeOf(input)
	fmt.Printf("%v Type is : %v\n", input, userType.Name())

	//获取值
	userValue := reflect.ValueOf(input)
	fmt.Printf("%v Value is : %v\n", input, userValue)

	//获取字段数量
	log.Println("user字段数量:", userType.NumField())

	//获取字段名和值
	log.Println("user字段名和值:")
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		value := userValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	log.Println("user方法数量:", reflect.TypeOf(input).NumMethod())

	//获取方法
	for j := 0; j < userType.NumMethod(); j++ {
		m := userType.Method(j) //对外开放的方法才能被获取到
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

//通过反射，重新赋值
func ReflectValue() {
	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num)

	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&num) //参数不仅可以是个interface，还可以是个指针
	newValue := pointer.Elem()       //获取指针指向的值

	fmt.Println("newValue:", newValue)
	fmt.Println("type of pointer:", newValue.Type())
	//通过CanSet方法查询是否可以设置返回false
	fmt.Println("settability of pointer:", newValue.CanSet())

	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)
}

func main() {
	user := User{1, "", 18}
	fmt.Printf("user: %v\n", user)
	user.SetName("hp")
	fmt.Printf("user: %v\n", user)

	//reflect的方法展示
	GetFieldAndMethod(user)

	//反射赋值
	ReflectValue()
}

//执行结果：
//user: {1  18}
//设置用户名称为:hp
//user: {1 hp 18}
//{1 hp 18} Type is : User
//{1 hp 18} Value is : {1 hp 18}
//2021/02/18 18:17:43 user字段数量: 3
//2021/02/18 18:17:43 user字段名和值:
//Id: int = 1
//Name: string = hp
//Age: int = 18
//2021/02/18 18:17:43 user方法数量: 1
//ReflectCallFunc: func(main.User)
