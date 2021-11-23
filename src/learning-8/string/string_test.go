package string

import "testing"

func TestString(t *testing.T) {

	var s string
	t.Log(s) // default ""
	t.Log(len(s))

	s = "hello"
	t.Log(s)
	t.Log(len(s))

	s = "\xE4\xB8\xA5" //可以存放任何二进制数据
	t.Log(s)
	t.Log(len(s))

	//s[3]= "s" 只读，不可以赋值

}

//1. string是数据类型， 而非引用/指针类型
//2. string是只读的byte 数组，len可以计算它所包含的byte数
//3. string的byte数组可以存放任何数据
