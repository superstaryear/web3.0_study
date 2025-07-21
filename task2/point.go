package main

import "fmt"

/*
指针题目1
编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func addTem(num *int) int {
	*num += 10
	return *num
}

/*
指针题目2
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func multSlices(nums *[]int) {
	for k, v := range *nums {
		(*nums)[k] = v * 2
	}
}

func main() {
	//指针题目1
	var num = 0
	fmt.Println(addTem(&num))
	//指针题目2
	nums := []int{1, 2, 3}
	fmt.Println("调用multSlices之前nums的值为：", nums)
	multSlices(&nums)
	fmt.Println("调用multSlices之后nums的值为：", nums)
}
