package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello world"

	//判断是不是以某个字符串开头，返回布尔值
	res0 := strings.HasPrefix(str, "http://")
	res1 := strings.HasPrefix(str, "hello")
	fmt.Printf("res0 is %v\n", res0)
	fmt.Printf("res1 is %v\n", res1)

	//判断是不是以某个字符串结尾
	res3 := strings.HasSuffix(str, "http://")
	res4 := strings.HasSuffix(str, "world")
	fmt.Printf("res3 is %v\n", res3)
	fmt.Printf("res4 is %v\n", res4)

	//判断字符在字符串中首次出现的索引位置，没有返回-1
	res5 := strings.Index(str, "o")
	res6 := strings.Index(str, "x")
	fmt.Printf("res5 is %v\n", res5) // 4
	fmt.Printf("res6 is %v\n", res6) // -1

	//返回字符最后一次出现的索引位置，没有返回-1
	res7 := strings.LastIndex(str, "o")
	res8 := strings.LastIndex(str, "x")
	fmt.Printf("res7 is %v\n", res7)  // 7
	fmt.Printf("res8 is %v\n", res8) // -1

	//字符串替换
	res9 := strings.Replace(str, "world", "golang", 2)
	res10 := strings.Replace(str, "world", "golang", 1)
	//trings.Replace("原字符串", "被替换的内容", "替换的内容", 替换次数)
	//原字符串中有2个world，才能替换2次
	fmt.Printf("res9 is %v\n", res9)  // res9 is hello golang
	fmt.Printf("res10 is %v\n", res10)  // res10 is hello golang

	//求字符在字符串中出现的次数，不存在返回0次
	countTime0 := strings.Count(str, "h")
	countTime1 := strings.Count(str, "x")
	fmt.Printf("countTime0 is %v\n", countTime0)
	fmt.Printf("countTime1 is %v\n", countTime1)

	//	重复几次字符串
	res11 := strings.Repeat(str, 0)
	res12 := strings.Repeat(str, 1)
	res13 := strings.Repeat(str, 2)
	// strings.Repeat("原字符串", 重复次数)
	fmt.Printf("res11 is %v\n", res11)  // res11 is
	fmt.Printf("res12 is %v\n", res12)  // res12 is hello world
	fmt.Printf("res13 is %v\n", res13)  // res13 is hello worldhello world

	//字符串改大写
	res14 := strings.ToUpper(str)
	fmt.Printf("res14 is %v\n", res14)

	// 字符串改小写
	res15 := strings.ToLower(str)
	fmt.Printf("res15 is %v\n", res15)

	// 去除首位的空格
	res16 := strings.TrimSpace(str)
	fmt.Printf("res16 is %v\n", res16)

	//去除首尾指定的字符,遍历l、d、e然后去除
	res17 := strings.Trim(str, "ld")
	fmt.Printf("res17 is %v\n", res17)  // res17 is hello wor

	// 去除开头指定的字符
	res18 := strings.TrimLeft(str, "he")
	fmt.Printf("res18 is %v\n", res18)

	//去除结尾指定的字符,遍历d、l、r
	res19 := strings.TrimRight(str, "dlr")
	fmt.Printf("res19 is %v\n", res19)

	// 用指定的字符串将string类型的切片元素结合
	str1 := []string{"hello", "world", "hello", "golang"}
	res20 := strings.Join(str1, "+")
	fmt.Printf("res20 is %v\n", res20)  // res20 is hello+world+hello+golang




}
