package main

import "fmt"

/*
变量是一种方便的占位符，用于引用计算机内存地址。

1.指针默认值nil
2.通过&(取地值符)取变量地址
3.通过*(取值符)透过指针访问目标值

首先基本数据类型中，如name="yugo" ，变量name存的值是yugo

1）基本数据类型，变量存的是值，称为值类型

2）通过&符号获取变量的地址，例如&name

3）指针类型的变量，存的是地址，这个地址指向的空间存的是值

4）获取指针类型指向的值，使用*，例如*ptr，使用*ptr获取ptr指向的值

5）值类型（int、float、bool、string、array、struct）都有对应指针类型，比如*int和*string等等
 */

func main() {
	var name string = "yugo"
	//查看name变量的内存地址，通过&name取地址
	fmt.Printf("name的内存地址：%v\n", &name)
	fmt.Printf("name的内存地址类型：%T\n", &name)

	//指针变量，存的是内存地址
	//ptr指针变量指向变量name的内存地址
	var ptr *string = &name
	fmt.Printf("指针变量ptr的内存地址：%v\n", ptr)

	//获取指针变量的值，用*ptr
	//*ptr表示读取指针指向变量的值
	fmt.Printf("指针变量ptr指向的值是：%v\n", *ptr)

}