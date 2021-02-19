package main

import "fmt"

type Rect struct {
	width  float64
	height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

//*Rect 指针类型
func (r *Rect) SetWidth(w float64) {
	//需要改变width的值,所有要通过改变Rect的内存地址
	r.width = w
}

func main() {
	//值类型
	//值类型包括基本数据类型，int,float,bool,string,以及数组和结构体(struct)。
	//值类型变量声明后，不管是否已经赋值，编译器为其分配内存，此时该值存储于栈上。
	//值类型的默认值：
	var a int                               //int类型默认值为 0
	var b string                            //string类型默认值为 nil空
	var c bool                              //bool类型默认值为false
	var d [2]int                            //数组默认值为[0 0]
	fmt.Println("a 变量的内存地址：", &a)           //a 变量的内存地址： 0xc00001a0a0  默认已经分配内存地址，可以使用&来取内存地址
	fmt.Println("a 的默认值：", a)               //a 的默认值： 0
	fmt.Println("b 的默认值：", b)               //b 的默认值：
	fmt.Println("c 的默认值：", c)               //c 的默认值： false
	fmt.Println("d 的默认值：", d)               //d 的默认值： [0 0]
	fmt.Println("React(结构体） 的默认值：", Rect{}) //React(结构体） 的默认值： {0 0}

	//引用类型
	//引用类型包括指针，slice切片，map ，chan，interface。
	//变量直接存放的就是一个内存地址值，这个地址值指向的空间存的才是值。所以修改其中一个，另外一个也会修改（同一个内存地址）。
	//引用类型必须申请内存才可以使用，make()是给引用类型申请内存空间。
	var ee []int
	//ee[1] = 100 //无法使用
	//引用类型需要申请内存才可以使用，make()是给引用类型申请内存的
	ee = make([]int, 5, 5)          //切片的初始化
	ee[1] = 100                     //申请内存后，可以使用
	fmt.Println("ee 的值：", ee)       //ee 的值： [0 100 0 0 0]
	fmt.Println("ee[1] 的值：", ee[1]) //ee[1] 的值： 100

	//声明的同时赋值，此时已经进行了初始化，不必进行make()操作
	var e = []int{1, 2, 3, 4, 5}
	fmt.Println("e 的值：", e) //e 的值： [1 2 3 4 5]
	e[2] = 200
	fmt.Println("e[2] 的值：", e[2]) //e[2] 的值： 200
	fmt.Println("e 的值：", e)       //e 的值： [1 2 200 4 5]

	f := make([]int, 5, 5)
	copy(f, e)             //copy是值类型，修改f的值，e的值不会改变
	fmt.Println("f的值：", f) //f的值： [1 2 200 4 5]
	f[1] = 100
	fmt.Println("f的值：", f) //f的值： [1 100 200 4 5]
	fmt.Println("e的值：", e) //e的值： [1 2 200 4 5]

	//*和&的区别 :
	//& 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
	//*是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .

	//求矩形的面积
	rect1 := Rect{10, 5}
	fmt.Println("Rect的面积是：", rect1.area()) //Rect的面积是： 50
	//设置rect的宽
	rect1.width = 5
	area := rect1.area()
	fmt.Println("Rect的值：", rect1)  //Rect的值： {5 5}
	fmt.Println("Rect的面积是：", area) //Rect的面积是： 25

	//修改Rect的宽度
	rect2 := &Rect{100, 200}
	rect2.SetWidth(3)
	fmt.Println("Rect2的面积：", rect2.area()) //Rect2的面积： 600
}
