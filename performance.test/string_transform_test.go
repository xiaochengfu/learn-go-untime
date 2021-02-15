package main

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)
func BenchmarkStringItoa(b *testing.B) {
	num := math.MaxInt64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
	b.StopTimer()
}

func BenchmarkStringFormat(b *testing.B) {
	var num int64 = math.MaxInt64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
	b.StopTimer()
}

func BenchmarkStringSprintf(b *testing.B) {
	var num int64 = math.MaxInt64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v", num)
	}
	b.StopTimer()
}

//结论：性能依次降低
//Format > Itoa > sprintf
//Format与Itoa性能几乎一样，推荐使用Itoa

//执行测试 go test -bench=.  -benchtime=3s
//goos: darwin
//goarch: amd64
//pkg: go-study/performance.test
//BenchmarkStringSprintf-8   	28729862	       123 ns/op
//BenchmarkStringItoa-8      	65379051	        54.9 ns/op
//BenchmarkStringFormat-8    	64259707	        54.7 ns/op
//PASS
//ok  	go-study/performance.test	10.889s

//tip: benchmark性能测试介绍
//1. 基准测试的代码文件必须以_test.go结尾
//2. 基准测试的函数必须以Benchmark开头，必须是可导出的
//3. 基准测试函数必须接受一个指向Benchmark类型的指针作为唯一参数
//4. 基准测试函数不能有返回值
//5. b.ResetTimer是重置计时器，这样可以避免for循环之前的初始化代码的干扰
//6. 最后的for循环很重要，被测试的代码要放到循环里
//7. b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能
//8. 使用 go test 命令，加上 -bench= 标记，接受一个表达式作为参数, .表示运行所有的基准测试
