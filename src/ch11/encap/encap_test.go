package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"1", "Tom", 18}
	e1 := Employee{Name: "Tom", Age: 18}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 20
	e2.Name = "Jerry"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e is %T", &e)
	t.Logf("e is %T", e2)
}

func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id: %s, Name: %s, Age: %d", e.Id, e.Name, e.Age)
}

//func (e *Employee) String() string {
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("Id: %s, Name: %s, Age: %d", e.Id, e.Name, e.Age)
//}

func TestStructOperations(t *testing.T) {
	e := &Employee{"1", "Tom", 20}
	t.Log(&e.Name)
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
