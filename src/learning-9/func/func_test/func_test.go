package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestMultiValues(t *testing.T) {
	a, _ := returnMultiValues() //_忽略第二个返回值
	t.Log(a)
}

//wrapper for inner function
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
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

//Var params
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParams(t *testing.T) {
	t.Log(Sum(1, 2, 3))
	t.Log(Sum(1, 2, 3, 5))
}

func Clear() {
	fmt.Println("Clear resources")

}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
	fmt.Println("End")
}

//函数 --- 一等公民
//1.可以有多个返回值
//2.所有的参数都是值传递
//3.函数可以作为变量的值
//4.函数可以作为参数和返回值

//计算函数执行时间
//Defer函数 类似Finally,在程序最后执行.用于释放资源释放锁
//Panic 程序中断，不会影响Defer
