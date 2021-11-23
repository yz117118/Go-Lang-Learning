package operation_char_test

import "testing"

func TestComparator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 4, 5, 6}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == b)
	//t.Log(a == c) 维数不同，不能比较，编译报错
	t.Log(a == d)
}

func Test(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 4, 5, 6}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == b)
	//t.Log(a == c) 维数不同，不能比较，编译报错
	t.Log(a == d)
}

const (
	// << 表状态
	Readable   = 1 << iota //iota = 0
	Writable               //iota = 1
	Executable             ////iota = 2
)

func TestBitClear(t *testing.T) {
	a := 7 //0111
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

//算术运算符
// 运算符 + - * / % ++ --
//比较运算符
// 运算符 == != < > >= <=
//注意: Go语言没有前置的＋＋，——
//	   数组只有维度相同，元素个数相等才能进行比较
//	   位运算 右边二进制位为1， 则结果为0，  右边为0， 看左边
