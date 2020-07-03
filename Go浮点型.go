package main

import "fmt"

/*
Go 语言提供了两种精度的浮点数，float32占用4个字节 和 float64占用8个字节，编译器默认声明为float64

小数类型就是存放小数的，如 1.2  0.005  -2.32

浮点数形式：浮点数=符号位+指数位+位数位

 */

func main() {
	//var price float32 = 100.02
	//fmt.Printf("price类型是：%T，值是%v", price, price)  // %T 类型  %v 默认值


	/*
	浮点数=符号位+指数位+位数位
	 */
	var price float32 = 11.22 //正数符号
	fmt.Println("price=", price)
	var num1 float32 = -3.4 //负数符号
	var num2 float64 = -8.23
	fmt.Println("num1=", num1, "num2=", num2)
	//尾数可能丢失，精度缺损
	var num3 float32 = -123.11111111105//精度丢失了
	var num4 float64 = -123.11111111105//float64的精度高于float32
	fmt.Println("num3=", num3, "num4=", num4)
	//输出结果
	//num3= -123.111115 num4= -123.11111111105
}
