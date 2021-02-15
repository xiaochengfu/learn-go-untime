package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func stringAdd(num int) string {
	numbers := 100
	var str string
	for i := 0; i < num; i++ {
		str = "" //保证与其他函数拼接效果一致
		for j := 0; j < numbers; j++ {
			str +=  strconv.Itoa(j)
		}
	}
	return str
}

func stringBuffer(num int) (str string) {
	numbers := 100
	for i := 0; i < num; i++ {
		var buffer bytes.Buffer
		for j := 0; j < numbers; j++ {
			buffer.WriteString(strconv.Itoa(j))
		}
		str = buffer.String()
	}
	return str
}

func stringBuilder(num int) (str string) {
	numbers := 100
	for i := 0; i < num; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		str = builder.String()
	}
	return str
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	stringAdd(b.N)
	b.StopTimer()
}

func BenchmarkStringBuffer(b *testing.B) {
	b.ResetTimer()
	stringBuffer(b.N)
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	stringBuilder(b.N)
	b.StopTimer()
}

func main() {
	//想要查看打印结果，可以将文件名改为man.go
	strAdd := stringAdd(100)
	fmt.Println(strAdd)
	fmt.Println("-----------------")
	strBuffer := stringBuffer(100)
	fmt.Println(strBuffer)
	fmt.Println("-----------------")
	strBuilder := stringBuilder(100)
	fmt.Println(strBuilder)
}

//结论：性能依次降低
//Builder > Buffer > add+
//推荐使用 Builder进行字符串拼接

//执行测试 go test -bench=.
//goos: darwin
//goarch: amd64
//pkg: go-study/performance.test
//BenchmarkStringAdd-8       	  204414	      5677 ns/op
//BenchmarkStringBuffer-8    	 1000000	      1093 ns/op
//BenchmarkStringBuilder-8   	 1330159	       902 ns/op