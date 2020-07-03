package main

import (
	"fmt"
	"strconv"
)

/*
Go基本数据类型转化

	Golang在不同的数据类型之间赋值时需要显示转化，无法自动转换。

	语法：

	T(v)  将v转化为T类型
	T  数据类型 如 int32 int64 float32
	V  需转化的变量
	类型转化注意点：

	1）Go中数据类型的转换可以从 数值范围大 > 数值范围小，反之也可，注意别溢出

	2）被转化的变量存储的数值，变量本身的数据类型没变化

	3）不同类型变量之间的计算

Go基本数据类型与string转化

	开发中经常将数据类型转成string

	方法1：
	fmt.Sprintf(“%参数”,表达式)

	方法2：
	fmt.Strconv

string类型转基本数据类型


 */

func main() {

	////2）被转化的变量存储的数值，变量本身的数据类型没变化
	//var num1 int32 = 100
	//var num2 float32 = float32(num1)  // num2强转为浮点型
	//fmt.Printf("num1=%v num2=%v\n", num1, num2)  // %v 值的默认格式
	//fmt.Printf("num1类型是：%T\n", num1)  // 本身的类型没有变化
	//
	////3）不同类型变量之间的计算
	//var n1 int32 = 12
	//var n2 int64 = 10
	////n3:=n1+n2 //不同类型之间无法计算,需要强转
	//n3 := int64(n1) + n2
	//fmt.Println(n3)

	////Go基本数据类型与string转化 方式1
	//var num1 int = 66
	//var num2 float64 = 25.25
	//var b bool = true
	//var myChar byte = 'c'
	////%q 单引号
	////%d 十进制表示
	//str1 := fmt.Sprintf("%d", num1)
	//fmt.Printf("str1 type %T str=%q\n", str1, str1)
	////%f 有小数点
	//str2 := fmt.Sprintf("%f", num2)
	//fmt.Printf("str2 type %T str2=%q\n", str2, str2)
	////%t 布尔值
	//str3 := fmt.Sprintf("%t", b)
	//fmt.Printf("str3 type %T str3=%q\n", str3, str3)
	////%c Unicode码对应的字符
	//str4 := fmt.Sprintf("%c", myChar)
	//fmt.Printf("str4 type %T str4=%q\n", str4, str4)

	////Go基本数据类型与string转化 方式2
	//var num1 int = 99
	//var num2 float64 = 66.66
	//var b1 bool = true
	//str1 := strconv.FormatInt(int64(num1), 10)
	//fmt.Printf("str1类型是%T str1=%q\n", str1, str1)  // str1类型是string str1="99"
	////参数解释
	//// f 格式
	//// 10 小数位保留10位
	//// 64  表示float64
	//str2 := strconv.FormatFloat(num2, 'f', 10, 64)
	//fmt.Printf("str2类型是%T str2=%q\n", str2, str2)  // str2类型是string str2="66.6600000000"
	//str3 := strconv.FormatBool(b1)
	//fmt.Printf("str3类型是%T str3=%q\n", str3, str3)  // str3类型是string str3="true"
	////Itoa，将int转为string
	//var num3 int64 = 1123
	//str4 := strconv.Itoa(int(num3))  // 必须强转int()
	//fmt.Printf("str4类型是%T str4=%q\n", str4, str4)  // str4类型是string str4="1123"

	//string类型转基本数据类型
	/*
	   ParseBool，ParseFloat，ParseInt和ParseUint将字符串转换为值
	   如转换失败，返回新类型默认值
	*/
	var str1 string = "true"
	var b1 bool
	/*
		strconv.ParseBool(str1)函数返回2个值（value bool, err error）
	 */
	b1, _ = strconv.ParseBool(str1)
	fmt.Printf("b 值类型= %T  b值=%v\n", b1, b1)  // b 值类型= bool  b值=true

	var str2 string = "1234"
	var num1 int64
	/*
		func ParseInt（s string，base int，bitSize int）（i int64，err error）
	 */
	num1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("num1类型：%T num1值：%v\n", num1, num1)  // num1类型：int64 num1值：1234

	var str3 string = "123.456"
	var float1 float64
	/*
		func ParseFloat（s string，bitSize int）（float64，error）
	 */
	float1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("float1类型：%T  float1值：%v", float1, float1)  // float1类型：float64  float1值：123.456





}