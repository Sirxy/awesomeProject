package main

/*
iota用于生成一组相似规则初始化的常量，在const常量声明的语句中，第一个常量所在行，iota为0，之后每一个常量声明加一。

例如time包的例子，一周7天，每天可以定义为常量，1~6，周日为0，这种类型也称为枚举
 */

import (
	"fmt"
)

const (
	Sunday = iota
	Monday  //通常省略后续行表达式
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)


//如果iota表达式被打断，需要显示恢复
const (
	A = iota //初始0
	B        // +1
	C = "c"  //iota枚举被打断 ，为  c
	D        // c，与上  相同。
	E = iota // 4，显式恢复。注意计数包含了 C、D 两个，此时为4 。
	F        // 恢复iota 加一，此时为5
)


func main() {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)  // 0 1 2 3 4 5 6
	fmt.Println(A, B, C, D, E, F)  // 0 1 c c 4 5

}