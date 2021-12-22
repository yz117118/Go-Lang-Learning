package oo2

import (
	"fmt"
	"testing"
)

type Code string

//1. interface define
type Programmer interface {
	WriteHelloWorld() Code
}

//2. implement interface， type
type GoProgrammer struct {
}

type JavaProgrammer struct {
}

//把方法绑定在一个结构上，方法名，返回值跟接口相同
func (p *GoProgrammer) WriteHelloWorld() Code {
	return "go is the best language"
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "Java is the best language"
}

//3.创建一个方法，我们可以传入不同的子类对同一命令实现不同的反应
func writeFirstProgram(p Programmer) {
	//T -> INSTANCE
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestInterface(t *testing.T) {
	//interface一定要传入子类指针实例
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
