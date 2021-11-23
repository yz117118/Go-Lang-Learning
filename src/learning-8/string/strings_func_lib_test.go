package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	t.Log(strings.Join(parts, "-"))
}

func TestStringConvert(t *testing.T) {
	//int to string
	s := strconv.Itoa(10)
	t.Log(s)
	//string to int
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}
