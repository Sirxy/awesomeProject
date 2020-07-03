package main

import (
	"fmt"
	"unsafe"
)

/*
Go布尔型

一个布尔类型的值只有两种:true 和 false。

if 和 for 语句的条件部分都是布尔类型的值，并且==和<等比较操作也会产生布尔型的值。
 */

func main() {

	var b = true
	fmt.Println("b=", b)  // b= true
	fmt.Println("b字节数=", unsafe.Sizeof(b))  // b字节数= 1



}
