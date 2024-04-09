package array

import "testing"

func TestArrayInit(t *testing.T) {
	// 创建初始化值为0
	var arr [3]int
	// 创建并初始化数组
	arr1 := [4]int{1, 2, 3, 4}
	// 不指定元素个数
	arr3 := [...]int{1, 2, 3, 5}
	arr1[1] = 5
	t.Log(arr[0], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	// 元素遍历，下划线为占位，表示有相应的返回结果，但并不关心
	for _, e := range arr3 {
		t.Log(e)
	}
}

/*
*
数组截取
*/
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 5}
	arr3_sec := arr3[:3]
	//arr3_sec1 := arr3[:-1] 不支持
	t.Log(arr3_sec)
}
