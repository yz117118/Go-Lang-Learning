package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	// << 表状态
	Readable   = 1 << iota //iota = 0
	Writable               //iota = 1
	Executable             ////iota = 2
)

func TestPrintConstantsTry1(t *testing.T) {
	t.Log(Monday, Wednesday)
}

func TestPrintConstantsTry2(t *testing.T) {
	a := 7 //0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	b := 1 //0001
	t.Log(b&Readable == Readable, b&Writable == Writable, b&Executable == Executable)
}
