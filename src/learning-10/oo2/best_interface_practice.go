package oo2

//接口使用原则
//1.接口定义尽量小，很多接口只有一个方法
//2.较大的接口定义，可以由多个小接口定义组成
//3.只依赖与必要功能的最小接口

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func StoreData(reader Reader) error {
	return nil
}
