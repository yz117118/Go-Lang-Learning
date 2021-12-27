package share_mem

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}

}

func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{} //需要事先知道有几个task
}

func cancel2(cancelChan chan struct{}) {
	close(cancelChan) //不用，且适用于多个接受者
}

func TestCancel(t *testing.T) {
	cancelChannel := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChannel)
	}
	//cancel1(cancelChannel)
	cancel2(cancelChannel)
	time.Sleep(time.Second * 1)
}
