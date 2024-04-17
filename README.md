# golang-gogogo
learning golang


字符串

与其他编程语言的差异
1. string 是数据类型，不是引用或指针类型
2. string 是只读的 byte slice，len 函数可以查看它所包含的 byte 数
3. string 的 byte 数组可以存放任何数据，包含完全不属于可见字符编码的也可存放

unicode 字符集
UTF8 是 Unicode的存储实现（转换为字符的规则）


函数：一等公民

与其他主要编程语言的差异
1. 可以有多个返回值
2. 所有参数都是值传递：slice, map, channel 会有传引用的错觉
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值

可变参数

不必指定参数的个数

defer 函数 类似于 finally

---

## 面向对象

### 封装数据和行为
结构体定义
```go
type Employee struct {
	Id int
	Name string
	Age int
}
```
实例创建及初始化
```go
e := Employee{Id: 1, Name: "Tom", Age: 18}
e1 := Employee{Name: "Tom", Age: 18}
e2 := new(Employee) // 返回的引用/指针，相当于 e := &Employee{}
e2.Id = 1 // 与其他编程语言的差异：通过实例的指针访问成员不需要使用 ->
e2.Name = "Tom"
e2.Age = 18
```

行为（方法）定义

与其他主要编程语言的差异：
```go
// 在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, Age: %d", e.Id, e.Name, e.Age)
}

// 通常为了避免内存拷贝，使用第二种定义方式
func (e *Employee) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, Age: %d", e.Id, e.Name, e.Age)
}
```

Go 接口

与其他主要编程语言的差异：
1. 接口为非入侵性，实现不依赖于接口定义
2. 所以接口的定义可以包含在接口使用者包内

接口变量
```go
var prog Coder = &GoProgrammer{}
// 类型
type GoProgrammer struct {
}
// 数据
&GoProgrammer{}
```

自定义类型
```go

```