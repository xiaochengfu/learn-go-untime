package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

//type EmployeeList []Employee
//
//func (list EmployeeList) Filter (fn func(a *Employee) bool) EmployeeList {
//	var tmp EmployeeList
//	for _, v := range list {
//		if fn(&v) {
//			tmp = append(tmp, v)
//		}
//	}
//	return tmp
//}

//上面注释的代码，可以通过go:generate动态生成

//go:generate genny -in=filter.go -out=gen-filter.go -pkg=main gen "Generic=Employee"
func filterEmployeeExample() {

	var list = EmployeeList{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
		{"Alice", 23, 5, 9000},
		{"Jack", 26, 0, 4000},
		{"Tom", 48, 9, 7500},
	}

	var filter EmployeeList

	filter = list.Filter(func(e *Employee) bool {
		return e.Age > 40
	})

	fmt.Println("----- Employee.Age > 40 ------")
	for _, e := range filter {
		fmt.Println(e)
	}

	filter = list.Filter(func(e *Employee) bool {
		return e.Salary <= 5000
	})

	fmt.Println("----- Employee.Salary <= 5000 ------")
	for _, e := range filter {
		fmt.Println(e)
	}
}

func main() {
	filterEmployeeExample()
}
