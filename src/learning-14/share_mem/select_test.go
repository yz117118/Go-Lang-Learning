package share_mem

import (
	"fmt"
	"testing"
	"time"
)

//利用select机制
//实现多渠道处理，超时判断

func insert() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func query() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func AsyncOperations() chan string {
	retCh := make(chan string, 1) //buffered channel
	go func() {
		ret := insert()
		fmt.Println("returned result")
		retCh <- ret //put get之后，才会继续执行，否则阻塞
		fmt.Println("service exited")
	}()
	return retCh
}

func TestTimeOut(t *testing.T) {
	select {
	case ret := <-AsyncOperations():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
		//default:
		//	t.Log("default")
	}
}
