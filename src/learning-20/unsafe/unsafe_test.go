package unsafe

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type Customer struct {
	Name string
	Age  int
}

func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i))
	t.Log(f)
}

type MyInt int

func TestConvert(t *testing.T) {
	//合理的类型转换
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}

func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer

	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		//写在buffer里，每次读确保从buffer中读到最新的值
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}

	readDataFn := func() {
		//从这块指针上取到一块buffer
		data := atomic.LoadPointer(&shareBufPtr)
		//打印数据
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()

	}
}
