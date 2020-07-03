package main

import "fmt"

/*

 */

func main() {

	//字符串是一个不可改变的字节序列。
	//Go string通常是用来包含人类可读的文本。
	//文本字符串通常被解释为采用 UTF8 编码的 Unicode 码点。
	//Go的字符串由单个字节连接起来。
	var city string = "我爱北京天安门"
	println(city)
	// city[0]='1' 错误，go的字符串不可变

	//识别转义符
	var story string = "妖怪\n放了我师父"
	fmt.Println(story)

	//反引号，以原生形式输出，包括特殊字符，防止注入攻击
	story2 := `
你看这个灯，它又大又亮
你好
我是银角大王吧，你吃了\n吗?
`
	fmt.Println(story2)

	//字符串拼接，识别空格
	str3 := "你好" + "  我是孙悟空"
	fmt.Println(str3)

	str4 := "hello world"
	//索引取出值的码值，格式化输出
	fmt.Printf("str4=%c\n",str4[1])
	//输出str3的长度
	fmt.Println(len(str4))

	//多行字符串拼接，注意 +   加号 写在上一行
	myname := "wo" + "shi" + "bei" + "jing" +
		"sha" + "he" + "yu"
	fmt.Println(myname)

	myname1 := "hello world"
	for _, ret := range myname1 {
		fmt.Printf("ret=%c\n", ret)
	}

	//Go修改字符串的方法
	myname2 := "hello world"
	m1 := []rune(myname2)  //转化为[]int32的切片,rune是int32的别名
	m1[4] = '皮'  //修改索引对应的值
	myname2 = string(m1)  //类型强转，rune转为string
	fmt.Println(myname2)











}
