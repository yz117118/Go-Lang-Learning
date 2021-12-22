package client

import (
	//need add to go path
	"go-workspace/go-learning/src/learning-12/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
	//t.Log(series.getFibonacciSerie(5))
}

//package是基本复用模块单元，以首字母大写表明可被包外代码访问
//代码的package可以和所在目录不一致
//同一目录里的Go代码的package要保持一致

//init method
//init在main之前執行
//不同包的Init按照包的导入的依赖关系决定执行顺序
//每个包可以有多个Init函数
//包的每个源文件也可以有多个init函数
