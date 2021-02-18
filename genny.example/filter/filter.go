package main

import (
	"github.com/cheekybits/genny/generic"
)

type Generic generic.Type

type GenericList []Generic

func (g GenericList) Filter(fn func(e *Generic) bool) GenericList {
	var tmp GenericList
	for _, a := range g {
		if fn(&a) {
			tmp = append(tmp, a)
		}
	}
	return tmp
}
