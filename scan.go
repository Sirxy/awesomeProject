package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//// 示例1：fmt.Scan
	//var name string
	//fmt.Println("请输入用户名：")
	//fmt.Scan(&name)
	//fmt.Println(name)

	//// 示例2：fmt.Scan
	//var name string
	//var age int
	//fmt.Println("请输入用户名")
	//// 当使用Scan时，会提示用户输入
	//// 用户输入完成之后，会得到两个值：count,用户输入了几个值；err，用输入错误则是错误信息
	//_, err := fmt.Scan(&name, &age)
	//fmt.Println(err)
	//if err == nil {
	//	fmt.Println(name, age)
	//} else {
	//	fmt.Println("用户输入数据错误", err)
	//}
	//特别说明：fmt.Scan 要求输入两个值，必须输入两个，否则他会一直等待。

	//// 示例3：fmt.Scanln
	//var name string
	//fmt.Print("请输入用户名：")
	////_, _ = fmt.Scanln(&name)
	//fmt.Scanln(&name)
	//fmt.Printf(name)

	//// 示例4：fmt.Scanln
	//var name string
	//var age int
	//fmt.Println("请输入用户名:")
	//// 当使用Scanln时，会提示用户输入
	//// 用户输入完成之后，会得到两个值：count,用户输入了几个值；err，用输入错误则是错误信息
	//count, err := fmt.Scanln(&name, &age)
	//fmt.Println(count, err)
	//fmt.Println(name, age)
	//// 特别说明：fmt.Scanln 等待回车。

	//示例5
	//var name string
	//var age int
	//fmt.Println("请输入用户名：")
	//_, _ = fmt.Scanf(("我叫%s 今年%d岁"), &name, &age)
	//fmt.Println(name, age)

	//问题：读取一整行
	//当使用ftm.Scan等功能时，如果输入一整行且期间存在空格，则无法获取整行。
	//想要读取一行数据，这个功能可以基于标准输入来进行实现。
	reader := bufio.NewReader(os.Stdin)
	// line，从stdin中读取一行的数据（字节集合 -> 转化成为字符串）
	// reader默认一次能4096个字节（4096/3)
	//    1. 一次性读完，isPrefix=false
	//       2. 先读一部分，isPrefix=true，再去读取isPrefix=false
	line, _, _ := reader.ReadLine()
	data := string(line)
	fmt.Println(data)
}