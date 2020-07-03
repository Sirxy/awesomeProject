package main

import "fmt"

/*
Golang 中没有专门的字符类型，如果要存储单个字符(字母)，一般使用 byte 来保存。

普通字符串就是一串固定长度的字符连接起来的字符序列。

Go 的字符串是由单个字节连接起来的。

也 就是说对于传统的字符串是由字符组成的，而 Go 的字符串不同，它是由字节组成的。

Go的字符用单引号表示

Go的字符串用双引号表示
 */

func main() {
	var c1 byte = 'a'
	var c2 byte = '2'  //字符的2
	//直接输出byte的值，也就是输出对应的字符的码值
	fmt.Println("c1=", c1)  // c1= 97
	fmt.Println("c2=", c2)  // c2= 50
	//输出字符的值，需要格式化输出
	fmt.Printf("c1值=%c c2值=%c\n", c1, c2)  // c1值=a c2值=2

	/*

	//Go变量保存的byte 对应码值ASCII表，范围在[0-1,a-z,A-Z…]
	//如果保存的字符对应码大于255，应该使用int而不是byte，否则overflows byte异常
	var c3 int = '皮' //正确
	var c4 byte = '皮' //overflows byte 报错

	Go语言默认字符编码UTF-8，统一规定

	Go字符的本质是一个整数，直接打印是UTF-8编码的码值

	给与变量赋值整数，按%c格式化输出，得到的是unicode字符

	var c4 int = 22269
	fmt.Printf("c4=%c\n", c4)
	//输出结果c4=国
	Go语言允许使用转义符号”\”

	Go语言字符类型允许计算，相当于整数运算，因为字符拥有对应的Unicode码

	*/


}