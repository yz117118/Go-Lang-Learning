package microkernel

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type FileCollector struct {
	evtReceiver EventReceiver
	agtCtx      context.Context
	stopChan    chan struct{}
	name        string
	content     string
}

func NewCollect(name string, content string) *FileCollector {
	return &FileCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

func (c *FileCollector) Init(evtReceiver EventReceiver) error {
	fmt.Println("initialize file collector", c.name)
	c.evtReceiver = evtReceiver
	return nil
}

func (c *FileCollector) Start(agtCtx context.Context) error {
	fmt.Println("start file collector", c.name)
	for {
		select {
		//停止的时候是 agt.Stop() 方法里有调用 agt.cancel() , 即发送了 agtCtx.Done()
		case <-agtCtx.Done(): //kernel 完成了，Agent调用cancel, ctx会Done， 会放一个空的到stop channel
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(time.Millisecond * 50) //每隔一段时间放一个Event进去
			c.evtReceiver.OnEvent(Event{c.name, c.content})
		}
	}
}

func (c *FileCollector) Stop() error {
	fmt.Println("stop collector", c.name)
	select {
	case <-c.stopChan: //stop channel中能取到数据时
		return nil
	case <-time.After(time.Second * 1):
		return errors.New("failed to stop for timeout")
	}
}

func (c *FileCollector) Destroy() error {
	fmt.Println(c.name, "released resources.")
	return nil
}

func TestAgent(t *testing.T) {
	agt := NewAgent(100)
	c1 := NewCollect("c1", "1")
	c2 := NewCollect("c2", "2")
	agt.RegisterCollectors("c1", c1)
	agt.RegisterCollectors("c2", c2)
	if err := agt.Start(); err != nil {
		fmt.Printf("start error %v\n", err)
	}
	//重复启动会失败
	fmt.Println(agt.Start())
	time.Sleep(time.Second * 1)
	agt.Stop()
	agt.Destroy()
}
