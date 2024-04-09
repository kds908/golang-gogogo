package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1: %d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2: %d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3: %d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	// 在访问的Key不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在
	if v, ok := m1[3]; ok {
		t.Logf("key3 is existing, value is %d", v)
	} else {
		t.Log("key3 is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 9}
	// 第一个参数，在数组中是index，在map中是key
	for k, v := range m1 {
		t.Log(k, v)
	}
}
