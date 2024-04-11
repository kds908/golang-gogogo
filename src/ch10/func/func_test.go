package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

/*
@Param 内部函数
@Return 内部函数
*/
func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	// 如果要忽略一个值，将值的变量设为 “_”
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpend(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("clear resources.")
}

func TestDefer(t *testing.T) {
	// 在函数执行完后执行，且不受异常影响
	defer Clear()
	fmt.Println("start...")
	//panic("err")
	// panic 后的不会被执行
	//fmt.Println("end !")
}
