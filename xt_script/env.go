package main

import (
	"fmt"
)

var varMap = map[string]objIntf{
	"println": newFunc(func (al []objIntf) objIntf {
		sl := make([]interface{}, 0, len(al))
		for _, a := range al {
			sl = append(sl, a.asStr())
		}
		fmt.Println(sl...)
		return newInt(0)
	}),
}
