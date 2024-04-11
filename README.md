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