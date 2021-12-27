package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 10; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
		}
	}
}

func TestObjPoolWithNoAvailableObj(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			//will throw time out err when get 11 obj
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
}

func TestObjPoolFully(t *testing.T) {
	pool := NewObjPool(10)
	if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
		//对象池满了会over flow
		t.Error(err)
	}
}

func TestObjPoolWithReleaseAfterUse(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			//get one release one
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
}
