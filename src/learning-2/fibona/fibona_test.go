package fibona

import (
	"testing"
)

func TestFibList(t *testing.T) {
	a := 1
	b := 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	tmp := a
	a = b
	b = tmp
	t.Log(a, b)
}

func TestSimpleExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}

// 变量赋值（类型推断）: 1. var a int =1
//		   2. var(a int =1
//				  b int =1)
//		   3. a:=1
//		   4. a = 1
// 交换变量表达式: a, b = b, a
