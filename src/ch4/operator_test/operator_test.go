package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// 长度不同的数组，编译不通过
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	t.Log(a)
	a = a &^ Executable
	t.Log(a)
	t.Log(a & Readable)
	t.Log(a & Writeable)
	t.Log(a & Executable)
	t.Log(Readable)
	t.Log(Writeable)
	t.Log(Executable)

	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
