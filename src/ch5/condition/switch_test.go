package condition

import "testing"

/**
 * 1. 条件表达式不限制为常量或整数
 * 2. 单个case中，可以出现多个结果选项，使用逗号分隔
 * 3. 与C语言等规则相反，go 不需要 break 来明确退出一个case
 * 4. 可以不设定 switch 之后的条件表达式，在此种情况下，整个switch结构与多个if...else...的逻辑作用等同
 */

// 多结果
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("is not 0-3")
		}
	}
}

// 类似 if else
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("is not 0-3")
		}
	}
}
