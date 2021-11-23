package string

import "testing"

func TestStringUnicode(t *testing.T) {

	var s string
	s = "中" //可以存放任何二进制数据
	t.Log(s)
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c))

	t.Logf("中 unicode %x", c[0]) //x 16进制
	t.Logf("中 utf8 %x", s)

}

func TestStringToRune(t *testing.T) {

	s := "入目无他人，满眼皆是你"
	for _, c := range s {
		t.Logf("%[1]c %[1]d %[1]x", c) //[1]都是对第一个参数
	}

}

//Unicode是一种字符集
//UTF-8是unicode的存储实现（转换为字节序列的规则）
