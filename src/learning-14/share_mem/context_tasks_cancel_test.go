package share_mem

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//Context
//根Context：通过context.Background 0创建
//子Context： context.WithCancel（parentContext）创建
//ctx，cancel ：= context.WithCancel（context.Background（）)
//前Context被取消时，基于他的子context都会被取消
//接收取消通知   <—ctx.Done()

func isCtxCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}

}

func TestCtxCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCtxCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
