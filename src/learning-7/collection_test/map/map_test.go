package _map_test

import "testing"

func TestMapInit(t *testing.T) {

	m1 := map[string]int{"one": 1, "two": 2}
	t.Log(m1["one"])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m2[4])
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestMapAccessNotExitKey(t *testing.T) {
	//不存在， 默认为0
	m1 := map[int]int{}
	t.Log(m1[1])

	//赋值0， 返回0
	m1[2] = 0
	t.Log(m1[2])

	//需要判断key是不是存在
	if v, ok := m1[2]; ok {
		t.Logf("key 2 is %d", v)
	} else {
		t.Log("key 2 is not existing")
	}

}

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2}
	for k, v := range m1 {
		t.Log(k, v)
	}

}

// key 是否存在， 存在 是否为空值
// map key不存在，不会返回nil， 而是返回0， 需要额外判断key是否存在在
