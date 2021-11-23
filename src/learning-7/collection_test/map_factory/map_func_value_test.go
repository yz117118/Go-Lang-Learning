package map_factory_test

import (
	"testing"
)

func TestMapWithFunValue(t *testing.T) {

	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))

}

// func methodName(Type param) returnType{}
//map的value可以使一个函数
