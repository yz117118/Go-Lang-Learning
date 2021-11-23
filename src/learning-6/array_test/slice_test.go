package array

import "testing"

func TestSliceInit(t *testing.T) {

	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 5}
	t.Log(len(s1), cap(s1))

	// init one slice, type, len is 3, cap is 5
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	// init value 0,0,0 depend len
	t.Log(s2[0], s2[1], s2[2])
	//index out of range [3] with length 3
	//t.Log(s2[3])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))

	//growing cap * 2 auto
	s2 = append(s2, 1, 2)
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 1, 2, 3)
	t.Log(len(s2), cap(s2))

}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aus", "Sep", "Oct", "Nov", "Dec"}
	Q1 := year[0:3]
	t.Log(Q1)
	t.Log(len(Q1), cap(Q1))
	Q3 := year[6:9]
	t.Log(Q3)
	t.Log(len(Q3), cap(Q3))
	Q3[1] = "UNKNOWM"
	t.Log(Q3)
	t.Log(year)

}

func TestSliceCompare(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aus", "Sep", "Oct", "Nov", "Dec"}
	Q1 := year[0:3]
	t.Log(Q1)
	t.Log(len(Q1), cap(Q1))
	Q3 := year[6:9]
	t.Log(Q3)
	t.Log(len(Q3), cap(Q3))
	Q3[1] = "UNKNOWM"
	t.Log(Q3)
	t.Log(year)

	//ERROR slice can only be compared to nil
	//if Q1 == Q3 {
	//
	//}
}

//Section : ptr(指针) len(元素个数) cap(内部数组容量), 内部数组是一个多slice共享的数据结构v
//if old_cap < len, new_cap = old_cap * 2; then array copy
//Slice[x: y] len = y - x, cap = cap - x
// array vs slice   [...]int{}  []int{}
// array不能扩容， slice可以，系数为2
// array可以比较，切片只能跟nil比较
