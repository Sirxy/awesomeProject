package main

import (
	"fmt"
	"unsafe"
)

/*
Go常量

常量代表只读的，不可修改的值，用const关键字定义。

如同用常量定义 “π”之类的常数。

常量如同变量一样，可以批量声明，或者一组相关的常量。

常量的计算都在编译期间完成，并非运行期间！减少运行时的工作。

未使用的常量不会引发编译错误。(这点和变量不一样哦~)
 */

//常量定义且赋值
const World string = "世界"
//多常量初始化
const x, y int = 1, 2
//常量类型推断，字符串类型
const s1 = "Hello golang"
//常量组
const (
	e = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
	b bool = true
)
//常量组，可以除了第一个外其他的常量右边的初始化表达式可以省略
//如果省略初始化表达式，默认使用前面常量的表达式
//与上一个常量相同
const (
	cc1 = 1
	cc2
	cc3
	cc4 = "c4444"
	cc5
)
//常量也可以定义函数的返回值
const (
	f1 = "abc"  //长度为3的字符串类型
	f2 = len(f1)  //返回长度的函数结果
	f3 = unsafe.Sizeof(f2)  //返回f2所占用的内存大小
)


func main() {
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
}