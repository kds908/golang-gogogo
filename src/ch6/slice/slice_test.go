package slice

import "testing"

/*
*

	 切片与数组
		1. 容量是否可伸缩
		2. 是否可以比较
*/
func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		// 放入程序超出空间时，扩容，s指向新的存储空间地址
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	q2 := year[3:6]
	t.Log(q2, len(q2), cap(q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknown"
	t.Log(q2)
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	//b := []int{1, 2, 3, 4}
	// 切片只能与 空nil比较
	//if a == b {
	//	t.Log("equal")
	//}
	if a == nil {
		t.Log("a == nil")
	}
}
