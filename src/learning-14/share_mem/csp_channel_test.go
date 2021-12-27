package share_mem

import (
	"fmt"
	"testing"
	"time"
)

//CSP vs.Actor
//和Actor的直接通讯不同，CSP模式则是通过Channel进行通讯的，更松耦合一些。
//Go中channel是有容量限制并且独立于处理Groutine，而如Erlang，Actor模式 中的mailbox容量是无限的，接收进程也总是被动地处理消息。

//channel producer consumer需要同时在线,get后才会执行后续操作
//buffered channel 只要容量未满，producer就可以放消息，非堵塞

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func AsyncService() chan string {
	retCh := make(chan string, 1) //buffered channel
	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret //put get之后，才会继续执行，否则阻塞
		fmt.Println("service exited")
	}()
	return retCh
}

func TestService(t *testing.T) {
	retCh := AsyncService() //running
	otherTask()             //running parallel
	fmt.Println(<-retCh)    //get
}
