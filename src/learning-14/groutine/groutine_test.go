package groutine

import (
	"fmt"
	"testing"
	"time"
)

//Thead vs. Groutine

//1．创时默认的stack的大小
//·JDK5 以后 Java Thread stack 默认为1M
//·Groutine的Stack 初始化大小为2K

//2．和KSE（Kernel Space Entity）的对应关系
//·Java Thread是1：1
//·Groutine是M：N 减少系统空间到用户空间切换的开销

//M Stystem Thread
//P Processor
//G Goroutine

//高并发
//如果一个协程一直没执行完，守护进程则会对携程栈标记，将该协程插入到队列末尾，先执行其他协程，不会堵塞processor
//当某个协程被系统中断，processor会移动到另一个系统线程中执行其他协程，当中断协程被唤醒，协程会携带信息插入到其中一个processor的协程执行队列里

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)

		//共享 10个10
		//go func() {
		//	fmt.Println(i)
		//}()
	}
	time.Sleep(time.Microsecond * 50)

}
