package main

import (
	"fmt"
	"math"
)

func main() {
	////7.翻转整数
	//num := code.Reverse(-321)
	//fmt.Println(num)

	//9.回文数
	//num := 232
	//res := code.MyIsPalindrome(num)
	//if res == true{
	//	fmt.Println("是回文数:",num)
	//}else {
	//	fmt.Println("不是回文数:",num)
	//}

	//num := 232
	//res := isPalindrome(num)
	//if res == true{
	//	fmt.Println("是回文数:",num)
	//}else {
	//	fmt.Println("不是回文数:",num)
	//}

	//13.罗马数字转整数
	//num := code.RomanToInt("III")
	//fmt.Println(num)

	//14.最长公共前缀
	//var strs = []string{"flower", "flow", "flight"}
	//res := code.LongestCommonPrefix(strs)
	//fmt.Println("prefix:", res)

	//20.有效的括号
	//res := code.IsValidMy1("([{}]])")
	//fmt.Println("res:",res)

	//21. 合并两个有序链表
	//list1 := new(code.ListNode)
	//list2 := new(code.ListNode)
	//code.DoSingleNode(list1, list2)
	//list := code.MergeTwoListTwoPoint(list1, list2)
	//list.Traverse()

	//26.删除有序数组中的重复项
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	////nums := []int{10, 3, 2, 1, 1, 1, 6, 2, 3, 3, 4, 4, 6}
	//res := code.RemoveDuplicatesSlice(nums)
	//fmt.Println("res:", res)

	//64.二进制求和
	//str := code.AddBinary("010","11")
	//fmt.Println("AddBinary: ",str)

}

// HasTwoDecimalPlaces 判断 f 是否只有2位小数
func HasTwoDecimalPlaces(f float64) bool {
	const epsilon = 1e-9 // 精度阈值
	multiplied := f * 100
	return math.Abs(multiplied-math.Round(multiplied)) < epsilon
}

const MaxNum = 100000

// 100的阶乘
func factorial(num int) {
	x := make([]int, MaxNum)
	x[0] = 1
	size := 1

	for i := 2; i <= num; i++ {
		size = multiply(x, size, i)

	}
	fmt.Println("---Factorial of", num, "is:")
	for i := size - 1; i >= 0; i-- {
		fmt.Print(x[i])
	}
	fmt.Println()
}

func multiply(x []int, size int, num int) int {
	carry := 0
	for i := 0; i < size; i++ {
		prod := x[i]*num + carry
		x[i] = prod % 10
		carry = prod / 10
	}

	for carry != 0 {
		x[size] = carry % 10
		carry /= 10
		size++
	}

	return size
}
