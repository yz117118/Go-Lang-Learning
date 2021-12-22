package error_handle

import (
	"container/list"
	//"container/list"
	"errors"
	"testing"
)

//1.没有异常机制

//2.error类型实现了error接口
//type error interface {
//	Error() string
//}

var EmptyError = errors.New("list size can not be 0")

func isEmpty(l *list.List) (bool, error) {
	length := l.Len()
	if length == 0 {
		return false, EmptyError
	}
	return true, nil
}

//3.通过errors.New快速创建错误实例
func TestErrorHandle(t *testing.T) {
	l := list.New()
	//l.PushFront(1)
	if b, err := isEmpty(l); err != nil {
		if err == EmptyError {
			t.Error(err)
		}
	} else {
		t.Log(b)
	}
}

//最佳实践  早发现早抛出
