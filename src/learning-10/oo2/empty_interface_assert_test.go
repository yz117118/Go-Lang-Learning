package oo2

import (
	"fmt"
	"testing"
)

//空接口可以表示任何类型
//利用断言实现对空接口的类型转换

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("String", i)
		return
	}
	fmt.Println("unsupported data type")
}

func DoSomethingWithSwitch(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("unsupported data type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething(nil)
	DoSomethingWithSwitch(10)
	DoSomethingWithSwitch("10")
	DoSomethingWithSwitch(nil)
}
