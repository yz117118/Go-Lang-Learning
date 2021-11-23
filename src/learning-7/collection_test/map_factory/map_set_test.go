package map_factory

import "testing"

func TestMapForSet(t *testing.T) {

	set := map[int]bool{}
	set[1] = true
	n := 3
	//1. check whether key exist 保证元素的唯一性
	if set[n] {
		t.Log("1 is existing")
	} else {
		t.Logf("%d is not existing", n)
	}

	//添加元素
	set[3] = true
	//元素个数
	t.Log(len(set))
	//删除元素
	delete(set, 1)
	n = 1
	if set[n] {
		t.Log("1 is existing")
	} else {
		t.Logf("%d is not existing", n)
	}

}
