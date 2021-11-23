package conditions_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log(a)
	}
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Env")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchNonEx(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Env")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("Unknown")
		}
	}
}

//支持变量赋值 if var declaration,conditions{}
//switch 不限制常量或者整数
//自动break
//单case下可以有多个结果
//switch后面可不跟表达式，等价于if else
