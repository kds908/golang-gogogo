package type_test

import "testing"

type MyInt int64

/*
*
只支持明式转换，不支持任何类型的隐式转换，不支持别名到原类型的隐式转换
*/
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

/*
*
支持指针类型，但指针不支持运算
*/
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// go 语言不支持指针运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

/*
*
字符串是值类型，默认初始化是空字符串而不是空
*/
func TestString(t *testing.T) {
	var s string
	// string 初始化为空字符串
	t.Log("*" + s + "*")
	// string 初始化长度为 0
	t.Log(len(s))
	// 默认对空字符串判断，而不需要对空(nil)判断
	if s == "" {

	}
}
