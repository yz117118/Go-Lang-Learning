package microkernel

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	Running         = 1
	Waiting         = 0
	WrongStateError = errors.New("the state of collector is incorrect")
)

type Event struct {
	Source  string //事件源
	Content string //时间内容
}

type EventReceiver interface {
	// OnEvent 可以将event传递给任何实现了Collector的collector
	OnEvent(evt Event)
}

func (agt *Agent) OnEvent(evt Event) {
	//接收事件
	agt.evtBuf <- evt

}

func NewAgent(sizeEvtBuf int) *Agent {
	agt := Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan Event, sizeEvtBuf),
		state:      Waiting,
	}

	return &agt
}

// Agent 注册中心
type Agent struct {
	//拓展点，注册进入Agent的Collector
	collectors map[string]Collector
	//Collector回传给Agent的事件
	evtBuf chan Event
	//task cancel
	cancel context.CancelFunc
	//context of cancelable
	ctx context.Context
	//agent的运行状态，hold住agent的状态
	state int
}

// Collector 收集器，每个收集器在初始化时会传入一个时间接收源，也就是Agent。 Agent会对统一注册在内的collector进行init，start，stop
type Collector interface {
	Init(evtReceiver EventReceiver) error
	// Start 不同的collector可以cancel event
	Start(agtCtx context.Context) error
	Stop() error
	// Destroy release resources
	Destroy() error
}

// RegisterCollectors 将collectors注册进Agent
func (agt *Agent) RegisterCollectors(name string, collector Collector) error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.collectors[name] = collector
	return collector.Init(agt)
}

type CollectorsError struct {
	CollectorErrors []error
}

func (ce CollectorsError) Error() string {
	var strs []string
	for _, err := range ce.CollectorErrors {
		strs = append(strs, err.Error())
	}
	return strings.Join(strs, ";")
}

// EventProcessGroutine 模拟的收集具体的业务事件。这里的事件是由各个 collector 上报的
func (agt *Agent) EventProcessGroutine() {
	var evtSeg [10]Event
	for {
		for i := 0; i < 10; i++ {
			select {
			case evtSeg[i] = <-agt.evtBuf: // 将 collector 收集的事件放到 evtBuf 中
			case <-agt.ctx.Done(): // 执行上下文完成，结束
				return
			}
		}
		fmt.Println(evtSeg) // 当 collector 收集的事件满 10 个，打印一次。
	}
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.state = Running
	agt.ctx, agt.cancel = context.WithCancel(context.Background()) // 来一个取消的上下文和取消函数
	go agt.EventProcessGroutine()                                  //收集事件
	return agt.StartCollectors()                                   //启动所有的Collector

}

func (agt *Agent) StartCollectors() error {
	var err error
	var errs CollectorsError
	var mutex sync.Mutex
	for name, collector := range agt.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mutex.Unlock()
			}()
			err = collector.Start(ctx)
			mutex.Lock()
			if err != nil {
				errs.CollectorErrors = append(
					errs.CollectorErrors,
					errors.New(name+":"+err.Error()),
				)
			}
		}(name, collector, agt.ctx)
	}
	if len(errs.CollectorErrors) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		return WrongStateError
	}
	agt.state = Waiting
	//cancel -> context.Done()
	agt.cancel()
	return agt.stopCollectors()
}

func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Stop(); err != nil {
			errs.CollectorErrors = append(
				errs.CollectorErrors,
				errors.New(name+":"+err.Error()),
			)
		}
	}
	if len(errs.CollectorErrors) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) Destroy() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	return agt.destroyCollectors()
}

func (agt *Agent) destroyCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Destroy(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors,
				errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorErrors) == 0 {
		return nil
	}
	return errs
}
