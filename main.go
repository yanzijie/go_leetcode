package main

import (
	"fmt"
	"github.com/yanzijie/go_leetcode/code"
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
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//nums := []int{10, 3, 2, 1, 1, 1, 6, 2, 3, 3, 4, 4, 6}
	res := code.RemoveDuplicatesSlice(nums)
	fmt.Println("res:", res)

	//64.二进制求和
	//str := code.AddBinary("010","11")
	//fmt.Println("AddBinary: ",str)

}
