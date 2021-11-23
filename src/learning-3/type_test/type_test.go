package type_test

import (
	"math"
	"testing"
)

type MyInt int64 //alias

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
	t.Log(math.MaxInt32) //预定义值
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a //指针类型 存放变量地址
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) // print type

}

func TestStr(t *testing.T) {
	var a string
	t.Log("*" + a + "*")
	t.Log(len(a))
}

//基本数据类型
//bool
//string 是值类型， 初始值是“” 而不是nil
//int int8 int16 int32 int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte  alias for uint8
//rune  alias for int32,represents a Unicode code point float32 float64
//complex64 complex128

//注意: Go不允许任何形式的隐式类型转换，只存在强制转换
//	支持指针类型 但不支持指针运算
