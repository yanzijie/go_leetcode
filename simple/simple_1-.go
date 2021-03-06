package simple

import (
	"fmt"
	"strconv"
)

// TwoSum /*****************************1.两数之和start******************************/
/* TwoSum
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出 和 为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
/*
暴力破解法 : 遍历每一个元素，与后面的元素相加，看看是否等于target
时间复杂度:O(n方)
空间复杂度:O(1)	空间大小固定，不随变量改变
*/
func TwoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums)-1; j++ {
			if nums[i]+nums[j+1] == target {
				res[0] = i
				res[1] = j + 1
			}
		}
	}
	return res
}

// TwoSumMap  TwoSumMap map法:
/*
1. 把下标和对应的值放入map中，下标为value，值为key。如数组[2, 7, 11, 15]，放入完成之后 {2:0, 7:1, 11:2, 15:3}
2. 在放入的过程中判断 target-nums[i] 是否等于map中的某个key，等于的话，这个key的value(下标), 和nums[i]的下标就是我们要找的下标值
时间复杂度 : O(n) n是数组的长度
空间复杂度 : O(n) 空间大小随着n的变化而变化
*/
func TwoSumMap(nums []int, target int) []int {
	res := make([]int, 2)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v, ok := m[target-nums[i]]
		if ok {
			res[0] = i
			res[1] = v
			return res
		} else {
			m[nums[i]] = i
		}
	}
	return res
}

//逐步对比法, 用target减去每一个元素，去和之后的元素对比
func ThreeSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		var tmp int
		tmp = nums[i]
		//和后面的每一个元素进行对比
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+tmp == target {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return nil
}

/*****************************1.两数之和end******************************/

// Reverse /***************************7. 整数反转star*****************************/
/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
示例 1:
输入: 123
输出: 321

 示例 2:
输入: -123
输出: -321

示例 3:
输入: 120
输出: 21
注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2的31次方, 2的31次方−1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/
/*
弹出余数解法
时间复杂度: o(n)  x有多长，遍历执行多少次
空间复杂度: o(1)  空间大小固定，不随变量改变
*/
func Reverse(x int) int {
	var res int = 0

	for x != 0 {
		//弹出x最后一位，并且移除x最后一位，从尾部开始添加到新的int数字里面，循环计算
		//int型， 0.几都属于0
		res = res*10 + x%10 //弹出x最后一位, 加到新数字里面
		x = x / 10          //移除x最后一位
	}

	//判断翻转之后的数字是否溢出, 翻转之后统一判断，减少运算次数, 在循环中一直判断的话，可能也会占用一定的内存
	if res < -2147483648 || res > 2147483647 {
		return 0
	}

	return res
}

/***************************7. 整数反转end*****************************/

// MyIsPalindrome /*****************************9.回文数start******************************/
//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//示例 1:
//输入: 121
//输出: true
//示例 2:
//输入: -121
//输出: false
//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//示例 3:
//输入: 10
//输出: false
//解释: 从右向左读, 为 01 。因此它不是一个回文数
//*/
func MyIsPalindrome(x int) bool {
	//负数翻转不会相同,直接pass
	if x < 0 {
		fmt.Println("x = 0, error")
		return false
	}

	//弹出最后一位数,每次都乘10,组成新int
	var num, tmpNum, oldNum int
	oldNum = x
	for x != 0 {
		tmpNum = x % 10 //弹出最后一位数
		num = num*10 + tmpNum
		x = x / 10
	}

	if num == oldNum {
		return true
	}

	return false
}

/*****************************9.回文数end******************************/

//昂，筠，之，齐，扬，溪, 舟，弈, 茗，宁

// RomanToNumber /*****************************13. 罗马数字转整数*****************************/
//例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，
//所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
//
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//
//左边代表的值比右边小就减，否则就加，最后一位总是加,貌似和上面的规律没啥关系
//小值在左边就减小值，在右边就加小值
//IV : -1+5 = 4
//VI : 5+1 = 6
//IXL : -1-10+50 = 39
//IVX : -1-5+10 = 4
//IXV : -1+10+5 = 14
//IXVV : -1+10+5+5 = 14
//大的在后 前面的 减
//小的在后 前面的 加
//最后一位总是加
func RomanToNumber(s string) int {
	var res int
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	if len(s) == 1 {
		return m[s[0]]
	}
	//IXVV
	//这里面不会加最后一位, 因为sum += 的是 s[i-1], s[i]并不累加进去
	for i := 1; i < len(s); i++ {
		// m[X] > m[I]
		if m[s[i]] > m[s[i-1]] {
			res -= m[s[i-1]]
		} else {
			res += m[s[i-1]]
		}
	}

	//加最后一位, 最后一位怎么都是加
	res += m[s[len(s)-1]]
	return res
}

/*****************************13. 罗马数字转整数*****************************/

// LongestCommonPrefix /*****************************14.最长公共前缀*****************************/
//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。s
//
//所有输入只包含小写字母 a-z 。
//设第一个字符串为公共前缀，然后用公共前缀依次比较之后的字符串，更新公共前缀
//公共长度一定小于等于最短长度
func LongestCommonPrefix(s []string) string {
	commonPrefix := s[0] //原始公共前缀
	for i := 1; i < len(s); i++ {
		ss := whoIsLonger(s[i], commonPrefix) //找到最短的那个元素
		index := 0                            //公共前缀的数量
		for index < len(ss) && s[i][index] == commonPrefix[index] {
			index++
		}
		//更新最长公共前缀
		commonPrefix = commonPrefix[:index]
	}

	return commonPrefix
}
func whoIsLonger(str1, str2 string) string {
	if len(str1) > len(str2) {
		return str2
	}
	return str1
}

/*****************************14.最长公共前缀*****************************/

// IsValid /*****************************20.有效的括号*****************************/
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

例如:
输入：s = "()"
输出：true

输入：s = "()[]{}"  "{[()]}"
输出：true

输入：s = "([)]"
输出：false

思路：
遇到左括号就入数组，遇到右括号就和数组的最后一个左括号对比，符合就继续，不符合就报错
如果符合，数组最后的左括号要出数组, 最后如果还有右括号，但是左括号数组没有元素了-报错，
*/
func IsValid(s string) bool { //这个是性能比较好的的算法
	if len(s) <= 1 {
		return false
	}
	if s[0] == ')' || s[0] == '}' || s[0] == ']' {
		return false
	}

	//存储左边括号的字符数组切片
	leftArr := make([]byte, 0, len(s))
	//左括号进数组
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			leftArr = append(leftArr, s[i]) //进数组
		} else {
			if len(leftArr) == 0 { //数组里面没有左括号了
				return false
			}
			switch s[i] {
			case ')':
				if leftArr[len(leftArr)-1] != '(' {
					return false
				}
				leftArr = leftArr[:len(leftArr)-1] //切片最后一位出队
			case ']':
				if leftArr[len(leftArr)-1] != '[' {
					return false
				}
				leftArr = leftArr[:len(leftArr)-1] //切片最后一位出队
			case '}':
				if leftArr[len(leftArr)-1] != '{' {
					return false
				}
				leftArr = leftArr[:len(leftArr)-1] //切片最后一位出队
			default:
				return false
			}
		}

	}

	if len(leftArr) == 0 {
		return true
	}
	return false
}

func IsValidMy(s string) bool {
	str := make([]string, 0, len(s)) //括号进入切片
	for i := 0; i < len(s); i++ {
		str = append(str, string(s[i]))
	}
	//遇到有括号进行判断, 看前一个元素是否为匹配的左括号
	for i := 1; i < len(str); i++ {
		if str[i] == ")" || str[i] == "]" || str[i] == "}" {
			switch str[i] {
			case ")":
				if str[i-1] == "(" {
					//匹配,当前元素和前一个元素从切片中移除
					str = append(str[:i-1], str[i+1:]...)
					//更新下标i，重头开始找，不然len的长度减小，i不减，前面的元素就都不会遍历
					i = 0
				}
			case "]":
				if str[i-1] == "[" {
					str = append(str[:i-1], str[i+1:]...)
					i = 0
				}
			case "}":
				if str[i-1] == "{" {
					str = append(str[:i-1], str[i+1:]...)
					i = 0
				}
			default:
				return false
			}
		}
	}
	//完成之后看看切片是否还有数据,有就报错
	if len(str) > 0 {
		return false
	}
	return true
}

func IsValidMy1(s string) bool {
	//括号放入一个切片
	ks := make([]string, 0, 10)
	for i := 0; i < len(s); i++ {
		ks = append(ks, string(s[i]))
	}
	//遍历切片
	for i := 0; i < len(ks); i++ {
		//如果第一个就是左括号，错
		//遇到右括号对比前一个元素是不是对应左括号
		//是把两个元素从切片中移除
		//重置下标重头遍历
		switch ks[i] {
		case ")":
			if i == 0 || ks[i-1] != "(" {
				return false
			}
			ks = append(ks[:i-1], ks[i+1:]...)
			i = 0
		case "]":
			if i == 0 || ks[i-1] != "[" {
				return false
			}
			ks = append(ks[:i-1], ks[i+1:]...)
			i = 0
		case "}":
			if i == 0 || ks[i-1] != "{" {
				return false
			}
			ks = append(ks[:i-1], ks[i+1:]...)
			i = 0
		default:
		}
	}

	//遍历完了,切片中还有数据则是错的
	if len(ks) > 0 {
		return false
	}
	return true
}

/*****************************20.有效的括号*****************************/

// MergeTwoListsMy /*****************************21. 合并两个有序链表*****************************/
/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
func MergeTwoListsMy(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tail.Next = l2
			return head.Next
		}

		if l2 == nil {
			tail.Next = l1
			return head.Next
		}

		if l1.data > l2.data {
			tail.Next = l2
			l2 = l2.Next
			tail = tail.Next
			tail.Next = nil
		} else {
			tail.Next = l1
			l1 = l1.Next
			tail = tail.Next
			tail.Next = nil
		}
	}
	return nil

}

// RemoveDuplicates *****************************26. 删除有序数组中的重复项*****************************
/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);
// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
  print(nums[i]);
}
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素
*/
func RemoveDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	fmt.Println("after:", nums[:slow+1])
	return slow + 1
}

/******************************26. 删除有序数组中的重复项******************************/

// AddBinary *****************************67. 二进制求和*****************************
/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);
// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
  print(nums[i]);
}
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素
*/
func AddBinary(a string, b string) string {
	res := ""

	carry := 0
	lenA, lenB := len(a), len(b)
	var n int
	if lenA > lenB {
		n = lenA
	} else {
		n = lenB
	}

	//'1'-'0' = ascii码 49-48 = 1(int32类型)
	//'0'-'0' = ascii码 48-48 = 0(int32类型)
	for i := 0; i < n; i++ {
		//如果字符串最后一位是 '1' carry += 1
		//如果字符串最后一位是 '0' carry += 0
		// 只有i<len()才进去取,防止下标越界
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		//2%2=0 1%2=1, 3%2=1
		res = strconv.Itoa(carry%2) + res
		// 2/2==1,意思就是逢2进1
		// 1/2或者0/2==0,就是不进1
		carry /= 2
	}
	//全部加完了,最后carry还>0, 说明还有进1,就在最前面补上一个1
	if carry > 0 {
		res = "1" + res
	}

	return res
}

/******************************67. 二进制求和******************************/
