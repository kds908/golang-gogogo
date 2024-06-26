package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello world"
	t.Log(len(s))
	s = "\xE4\xBB\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))

	// rune 可以取出字符串字节切片
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConvert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("string:" + s)
	// 先判断转换无错误后再执行加值
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(100 + i)
	}
}
