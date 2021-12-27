package unit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//内置单元测试框架
//Fail, Error:该测试失败，该测试继续，其他测试继续执行
//FailNow, Fatal:该测试失败，该测试终止，其他测试继续执行

//go test -v -cover
//断言 go get -u github.com/stretchr/testify

//单元格测试法
func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("expected output is %d, actually output is %d", expected[i], ret)
		}
		t.Logf("expected output is %d, actually output is %d", expected[i], ret)
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("Start")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Fatal")
	fmt.Println("Start")
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 8}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}
