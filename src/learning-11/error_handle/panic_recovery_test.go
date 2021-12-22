package error_handle

import (
	"errors"
	"fmt"
	"testing"
)

//panic用于不可恢复的错误
//os.Exit 退出不会调用defer和打印堆栈信息
//recovery 错误回复可能会形成僵尸服务进程，导致health check 失败
//Let it Crash!往往是回复不确定错误的最好方法

func TestPanicVsExit(t *testing.T) {

	defer func() {
		fmt.Println("Finally!")
	}()
	fmt.Println("Start!")
	panic(errors.New("Something wrong"))
	//os.Exit(1)

}

func TestRecovery(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovery from", err)
		}
	}()

	fmt.Println("Start!")
	panic(errors.New("Something wrong"))

}
