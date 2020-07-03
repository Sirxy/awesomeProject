package main

import "fmt"

/*
数组是固定长度的特定类型元素组成的序列。

一个数字由零或多个元素组成。

数组的长度是固定，因此Go更常用Slice（切片，动态增长或收缩序列）。

数组是值类型，用索引下标访问每个元素，范围是0~len(数组)-1，访问越界会panic异常

赋值和传参是复制整个数组而不是指针


 */

func main() {

	//var a [5]int
	//a[0] = 1
	//a[1] = 2
	//fmt.Println(a)
	///*
	//   定义数组
	//   var 数组名 [数组大小]数据类型
	//   var a1 [5]int
	//   定义数组后，5个元素都有默认值 0
	//   数组赋值方式
	//   a[0]=1
	//   a[1]=2
	//   数组的第一个元素的地址，就是数组的首地址
	//   数组各个元素地址间隔根据数组的数据类型决定,int64 8字节  int32 4字节
	//*/
	//var intArr [5]int
	//fmt.Println("intArr默认值是：", intArr)
	//intArr[0] = 1
	//intArr[1] = 2
	//intArr[2] = 3
	//fmt.Println("intArr赋值后的值是：", intArr)
	//fmt.Printf("intArr数组地址是=%p\n", &intArr)
	//fmt.Printf("intArr数组第一个元素地址是=%p\n", &intArr[0])
	//fmt.Printf("intArr数组第二个元素地址是=%p\n", &intArr[1])
	//fmt.Printf("intArr数组第三个元素地址是=%p\n", &intArr[2])

	////(全局声明)
	////声明赋值方式
	//var a1 [5]string = [5]string{"大猫", "二狗"}
	//fmt.Println(a1)
	////自动类型推导，未赋值的有默认值
	//var a2 = [5]int{1, 2, 3}
	//fmt.Println(a2)
	////自动判断数组长度
	//var a3 = [...]int{1, 2, 3, 4, 5}
	//fmt.Println(a3)
	////指定索引赋值元素
	//var a4 = [...]string{3: "狗蛋", 6: "貌似"}
	//fmt.Println(a4)
	////结构体类型数组
	//var a5 = [...]struct {
	//	name string
	//	age int
	//}{
	//	{"王麻子", 10},
	//	{"吕秀才", 19},
	//}
	//fmt.Println(a5)


	/*
	遍历数组
	 */
	//var b1 = [...]int{1, 2, 3, 4, 5, 6}
	////通过索引取值
	//for i := 0; i < len(b1); i++ {
	//	fmt.Println(b1[i])
	//}
	////for循环遍历数组，索引和值，index可以省略用占位符_
	//for index, value := range b1 {
	//	fmt.Println(index, value)
	//}


	/*
	数组使用细节
	 */
	////数组是多个相同类型数据的组合，且长度固定，无法扩容
	//var a1 [3]int
	//a1[0] = 1
	//a1[1] = 11
	////必须赋值int类型数据，否则报错
	////a1[2] = 111.1
	////不得超出索引
	////a1[3]=111
	//fmt.Println(a1)//有默认值[1 11 0]


	/*
	数组使用步骤：

	声明数组
	给数组元素赋值
	使用数组
	数组索引从0开始，且不得越界否则panic
	Go数组是值类型，变量传递默认是值传递，因此会进行值拷贝
	修改原本的数组，可以使用引用传递（指针）
	 */
	//声明arr数组，需要考虑传递函数参数时，数组的长度一致性
	arr := [3]int{11, 22, 33}
	//test函数不会修改数组
	test(arr)
	fmt.Println(arr)
	//test2函数修改了数组
	test2(&arr)
	fmt.Println(arr)

}

//函数接收值类型，默认有值拷贝
func test(arr [3]int) {
	arr[0] = 66
}

//函数修改原本数组，需要使用引用传递
func test2(arr *[3]int) {
	//(*arr)[0] = 66 // 可以缩写arr[0]=66 编译器自动识别，arr是指针类型
	arr[0] = 66 // 可以缩写arr[0]=66 编译器自动识别，arr是指针类型
}
