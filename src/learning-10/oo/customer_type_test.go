package oo

import (
	"fmt"
	"testing"
	"time"
)

//alias of func(op int) int
//自定义类型
type IntConv func(op int) int
type IntConvSimple int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFnc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc(t *testing.T) {
	tsSF := timeSpent(slowFnc)
	t.Log(tsSF(10))
}
