package main

func main() {
	/*
	条件语句：最基本
	 */

	//if true {
	//	fmt.Println("666")
	//} else {
	//	fmt.Println("999")
	//}

	//if 1 > 2 {
	//	fmt.Println("666")
	//} else {
	//	fmt.Println("999")
	//}

	//flag := false
	//if flag {
	//	fmt.Println("条件成立")
	//} else {
	//	fmt.Println("条件不成立")
	//}

	//// 练习题1:用户输入姓名，判断是否正确
	//var name string
	//fmt.Print("请输入姓名:")
	//fmt.Scanln(&name)
	//if name == "alex" {
	//	fmt.Println("用户名输入正确")
	//} else {
	//	fmt.Println("用户名输入错误")
	//}

	//// 练习题2:用户输入数字，判断奇数、偶数
	//var number int
	//fmt.Println("请输入数字：")
	//fmt.Scanln(&number)
	//if number % 2 == 0 {
	//	fmt.Println("您输入的是偶数")
	//} else {
	//	fmt.Println("您输入的是奇数")
	//}

	////练习题3:用户和密码，判断用户名密码是否正确。
	////var username, password string
	//var (
	//	username string
	//	password string
	//)
	//fmt.Print("请输入用户名：")
	//fmt.Scanln(&username)
	//fmt.Print("请输入密码：")
	//fmt.Scanln(&password)
	//if username == "alex" && password == "sb" {
	//	fmt.Println("欢迎登陆pornhub")
	//} else {
	//	fmt.Println("用户名或密码错误")
	//}

	//// 练习题4:请输入用户名校验是否是VIP
	//var username string
	//fmt.Print("请输入用户名：")
	//fmt.Scanln(&username)
	//if username == "alex" || username == "eric" {
	//	fmt.Println("天上人间大VIP")
	//} else {
	//	fmt.Println("屌丝")
	//}


	/*
	条件语句：多条件判断
	 */
	//var length int
	//fmt.Println("请输入你的长度:")
	//fmt.Scanln(&length)
	//if length < 1 {
	//	fmt.Println("没用的东西，还特么是坑")
	//} else if length < 6 {
	//	fmt.Println("刚刚能用")
	//} else if length < 18 {
	//	fmt.Println("生活和弦")
	//} else {
	//	fmt.Println("太特么大了")
	//}


	/*
	条件语句：嵌套
	 */
	//fmt.Println("欢迎致电10086，1.话费相关；2.业务办理；3.人工服务。")
	//var number int
	//fmt.Scanln(&number)
	//if number == 1 {
	//	fmt.Println("话费服务，1.交话费；2.查询。")
	//	var n1 int
	//	fmt.Scanln(&n1)
	//	if n1 == 1 {
	//		fmt.Println("缴话费拉")
	//	} else if n1 == 2 {
	//		fmt.Println("查话费了")
	//	} else {
	//		fmt.Println("输入错误")
	//	}
	//} else if number == 2 {
	//	fmt.Println("业务办理")
	//} else if number == 3 {
	//	fmt.Println("人工服务")
	//} else {
	//	fmt.Println("输入错误")
	//}
	//// 建议：条件的嵌套不要太多


	/*
	今日作业
	 */

	////1.提⽰⽤户输入⿇花藤. 判断⽤户输入的对不对。如果对, 提⽰真聪明, 如果不对, 提⽰你 是傻逼么。
	//var name string
	//fmt.Print("请输入⿇花藤：")
	//fmt.Scanln(&name)
	//if name == "⿇花藤" {
	//	fmt.Println("真聪明")
	//} else {
	//	fmt.Println("是傻逼么")
	//}

	////2.提示用户输入两个数字，计算两个数的和并输出。
	//var (
	//	number1 int
	//	number2 int
	//)
	//var sum int
	//fmt.Println("请输入数字1：")
	//fmt.Scanln(&number1)
	//fmt.Println("请输入数字2：")
	//fmt.Scanln(&number2)
	//sum = number1 + number2
	//fmt.Printf("两个数的和为：%d", sum)

	////3.提示用户输入姓名、位置、行为三个值，然后做字符串的拼接 并输出，例如：xx 在 xx 做 xx 。
	//var username string
	//var location string
	//var action string
	//fmt.Print("请输入姓名：")
	//fmt.Scan(&username)
	//fmt.Print("请输入位置：")
	//fmt.Scan(&location)
	//fmt.Print("请输入行为：")
	//fmt.Scan(&action)
	//fmt.Printf("%s在%s做%s.", username, location, action)

	////4.设定一个理想数字比如：66，让用户输入数字，如果比66大，则显示猜测的结果大了；如果比66小，则显示猜测的结果小了;只有等于66，显示猜测结果正确。
	//const number int = 66
	//var input_number int
	//fmt.Println("请输入数字：")
	//fmt.Scan(&input_number)
	//if input_number > number {
	//fmt.Printf("你输入的数字为%d,比%d大。", input_number, number)
	//} else if input_number < number {
	//	fmt.Printf("你输入的数字为%d,比%d小。", input_number, number)
	//} else {
	//	fmt.Println("恭喜猜对了")
	//}

	////5.写程序，输出成绩等级。成绩有ABCDE5个等级，与分数的对应关系如下.
	////	A    90-100
	////	B    80-89
	////	C    60-79
	////	D    40-59
	////	E    0-39
	////	要求用户输入0-100的数字后，正确输出他的对应成绩等级。
	//var score int
	//fmt.Println("请输入您的分数：")
	//fmt.Scan(&score)
	// if 90 <= score && score <= 100 {
	// 	fmt.Println("您的成绩等级为A")
	// } else if 80 < score && score <= 89 {
	// 	fmt.Println("您的成绩等级为B")
	// } else if 60 <= score && score <= 79 {
	// 	fmt.Println("您的成绩等级为C")
	// } else if 40 <= score && score <= 59 {
	// 	fmt.Println("您的成绩等级为D")
	// } else if 0 <= score && score <= 39 {
	// 	fmt.Println("您的成绩等级为E")
	// } else {
	// 	fmt.Println("您输入的成绩有误")
	// }

}