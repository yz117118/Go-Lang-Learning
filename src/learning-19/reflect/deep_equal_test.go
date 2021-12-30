package reflect

import (
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two"}
	b := map[int]string{1: "one", 2: "two"}
	//t.Log(a == b)
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3", reflect.DeepEqual(s1, s3))

}
