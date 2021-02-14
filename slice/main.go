package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {

	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	fmt.Println(foo) //[0 0 0 42 100]

	bar := foo[1:4]
	bar[1] = 99
	fmt.Println(bar) //	[0 99 42]
	fmt.Println(foo) // [0 0 99 42 100]
	//结论：
	// 切片是内存共享的，foo改变，foo的切片bar也会改变

	a := make([]int, 32)
	b := a[1:16]
	a[2] = 42
	fmt.Println(a, len(a), cap(a), b, len(b), cap(b)) //[0 0 42 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 42 0 0 0 0 0 0 0 0 0 0 0 0 0]
	a = append(a, 1)
	a[2] = 43
	fmt.Println(a, len(a), cap(a), b, len(b), cap(b)) //[0 0 43 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1] [0 42 0 0 0 0 0 0 0 0 0 0 0 0 0]

	//结论
	//append()操作让 a 的容量变成了 64，而长度是 33。
	//这里你需要重点注意一下，append()这个函数在 cap 不够用的时候，就会重新分配内存以扩大容量，
	//如果够用，就不会重新分配内存了

	path := []byte("AAAA/BBBBBBBBB")
	fmt.Println(reflect.TypeOf(path))
	sepIndex := bytes.IndexByte(path,'/')
	fmt.Println(sepIndex)

	//dir1 := path[:sepIndex]
	dir1 := path[:sepIndex:sepIndex] //Full Slice Expression
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>",string(dir1),"dir1.len =>", len(dir1),"dir1.cap =>", cap(dir1)) //dir1 => AAAA dir1.len => 4 dir1.cap => 4
	fmt.Println("dir2 =>",string(dir2),"dir2.len =>", len(dir2),"dir2.cap =>", cap(dir2)) //dir2 => BBBBBBBBB dir2.len => 9 dir2.cap => 9

	dir1 = append(dir1,"suffix"...)

	fmt.Println("dir1 =>",string(dir1),"dir1.len =>", len(dir1),"dir1.cap =>", cap(dir1)) //dir1 => AAAAsuffix dir1.len => 10 dir1.cap => 16
	fmt.Println("dir2 =>",string(dir2),"dir2.len =>", len(dir2),"dir2.cap =>", cap(dir2)) //dir2 => BBBBBBBBB dir2.len => 9 dir2.cap => 9

	//新的代码使用了 Full Slice Expression，最后一个参数叫“Limited Capacity”，于是，后续的 append() 操作会导致重新分配内存。


}
