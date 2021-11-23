package array_test

import "testing"

func TestArrayInt(t *testing.T) {

	var arr1 [3]int
	arr1[0] = 1

	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4}

	t.Log(arr1, arr2, arr3)

}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	//normal loop
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	//for each index, value
	for idx, value := range arr3 {
		t.Log(idx, value)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 5}
	arr3Sec := arr3[3:]
	t.Log(arr3Sec)
}

func TestArrayCompare(t *testing.T) {
	//arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4}
	arr4 := [...]int{1, 2, 3, 5}
	arr5 := [...]int{1, 2, 3, 4}
	//len should be same
	//if arr3==arr2 {
	//	t.Log("true")
	//}
	if arr3 == arr4 {
		t.Log("true")
	} else if arr3 == arr5 {
		t.Log("true")
	}

	// array compare
	//1. 长度相等
	//2. 元素相等

}
