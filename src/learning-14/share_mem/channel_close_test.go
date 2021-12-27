package share_mem

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

//问题1， receiver不知道producer会放几个ch
//问题2. 放一个标记位-1，收到-1则结束，不适用于多个producer

//channel的关闭
//向关闭的channel 发送数据，会导致panic
//v，ok ＜-ch；ok为bool值，true表示正常接受，false 表示通道关闭
//所有的 channel 接收者都会在channel关闭时，立刻从阻塞等待中返回且上述ok 值为 false。这个广播机制常被利用，进行向多个订阅者同时发送信号。
