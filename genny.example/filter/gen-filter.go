// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

type EmployeeList []Employee

func (g EmployeeList) Filter(fn func(e *Employee) bool) EmployeeList {
	var tmp EmployeeList
	for _, a := range g {
		if fn(&a) {
			tmp = append(tmp, a)
		}
	}
	return tmp
}
