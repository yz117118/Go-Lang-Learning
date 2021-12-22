package oo

import "testing"

type Code struct {
	code string
}

//1. interface define
type Programmer interface {
	WriteHelloWorld() Code
}

//2. implement interface， type
type GoProgrammer struct {
}

//把方法绑定在一个结构上，方法名，返回值跟接口相同
func (p *GoProgrammer) WriteHelloWorld() Code {
	return Code{"hello world"}
}

func TestInterface(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Logf("p is %T", p) //p point
	t.Log(p.WriteHelloWorld())
}

//1.接口非入侵性，不用显示实现，实现不依赖于接口定义 //Duck type
//2.接口的定义可以包含在接口使用者包内，不用单独创建一个只有接口的包

//接口变量的两部分: type + data
//var p Programmer = &GoProgrammer{}
//type: Goprogrammer data: &GoProgrammer{}
