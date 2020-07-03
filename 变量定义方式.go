package main

import "fmt"

/*
变量定义方式
 */

//一次性声明多个全局变量

//声明全局变量方式1
var m1 = 100
var m2 = 200
var m3 = 300

//声明全局变量方式2
var (
	d1, d2, d3 = 1, 2, 3
)

//声明全局变量方式3
var (
	c1 = 100
	c2 = 200
	c3 = 300
)


func main() {

	//var name string
	//name = "超哥"
	//fmt.Println("name的值是：", name)
	//
	////定义变量与类型，不赋值，含有默认值
	//var salary float64
	//fmt.Println("salary的值是：", salary)
	//
	////一次性定义多个变量
	//var num, num1 = 10, 11
	//fmt.Println(num, num1)
	//
	////短声明，省略var，只能用在函数内部
	//username := "sirxy"
	//fmt.Println(username)
	//
	////一次性声明多个变量，int默认值
	//var n1, n2, n3 int
	//fmt.Println(n1, n2, n3)
	//
	////声明多个变量，且赋值
	//var c1, c2, c3 = "sirxy", 18, 99.99
	//fmt.Println(c1, c2, c3)
	//
	////短声明多个变量
	//a1, a2, a3 := "sirxy", 19, 100.0
	//fmt.Println(a1, a2, a3)
	//
	////特殊变量，占位符 “_”
	//_, name = Person(18, "好海奥")
	//fmt.Println(name)

	//常见数据类型变量默认值
	// 只声明变量，不赋值，只有默认值
	var age int
	var name string
	var gender bool
	var salary float64
	fmt.Println("age默认值 :", age)
	fmt.Println("name默认值 :", name)
	fmt.Println("gender默认值 :", gender)
	fmt.Println("salary默认值 :", salary)
}

func Person(a1 int, n1 string) (int, string) {
	return a1, n1
}
