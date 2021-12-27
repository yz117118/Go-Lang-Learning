package obj_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

//sync.Pool 只能当作一个缓存，而不能作为对象池
//适合于通过复用，降低复杂对象的创建和GC代价
//协程安全,会有锁的开销
//生命周期受GC影响，不适合于做连接池等，需要自己管理生命周期的资源的池化

//Processor :   私有对象	 协程安全
//				共享池    协程不安全

//生命周期
//GC会clear sync.pool缓存的对象
//对象的缓存有效期在下一次GC之前

//Get : 尝试从私有对象获取
//		私有对象不存在，尝试从当前Processor的共享池获取
//		如果当前Processor共享池也是空的，尝试从别的Processor的共享池中获取
//		如果所有子池为空，最后就用用户指定的New函数产生一个新的对象返回

//Put : 如果私有对象不存在，则保存为私有对象
//		如果私有对象已经存在，则往当前P的共享池里面放

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object")
			return 100
		},
	}
	v := pool.Get().(int) //第一次取出，Processor的私有对象为空，创建为私有对象
	fmt.Println(v)
	pool.Put(3)               //放回
	v1, _ := pool.Get().(int) //不会再创建私有对象了
	fmt.Println(v1)

}

func TestSyncPoolWithGC(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object")
			return 100
		},
	}
	v := pool.Get().(int) //第一次取出，Processor的私有对象为空，创建为私有对象
	fmt.Println(v)
	pool.Put(3) //放回
	runtime.GC()
	//3会被清空，下面再Get，会重新调用创建方法
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolWithGC2(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object")
			return 100
		},
	}
	v := pool.Get().(int) //第一次取出，Processor的私有对象为空，创建为私有对象
	fmt.Println(v)
	pool.Put(3) //放回
	runtime.GC()
	//3会被清空，下面再Get，下面再Get，会重新调用创建方法
	v1, _ := pool.Get().(int) //不会再创建私有对象了
	fmt.Println(v1)
	v2, _ := pool.Get().(int) //取出对象，私有对象池没有了
	fmt.Println(v2)
}

func TestSyncPoolWithMultiGrotine(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			//先取出sync pool中存在的对象，当拿不到已存在对象时，才回去创建新的对象取出
			fmt.Println(pool.Get().(int))
			wg.Done()
		}(i)
	}
	wg.Wait()

}
