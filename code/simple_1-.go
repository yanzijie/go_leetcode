package code

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
	m := make(map[int]int)
	for i, aLen := 0, len(nums); i < aLen; i++ {
		_, ok := m[target-nums[i]]
		if ok {
			return []int{m[target-nums[i]], i}
		}
		m[nums[i]] = i
	}
	return nil
}

// ThreeSum 逐步对比法, 用target减去每一个元素，去和之后的元素对比
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
示例1:
输入: 123
输出: 321

示例 2:
输入: -123
输出: -321

示例 3:
输入: 120
输出: 21
注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为[−2的31次方,2的31次方−1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
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
// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 示例 1:
// 输入: 121
// 输出: true
// 示例2:
// 输入: -121
// 输出: false
// 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
// 示例 3:
// 输入: 10
// 输出: false
// 解释: 从右向左读, 为 01 。因此它不是一个回文数
// */
func MyIsPalindrome(x int) bool {
	// 负数翻转不会相同,直接pass
	if x < 0 {
		return false
	}

	//弹出最后一位数,每次都乘10,组成新int
	var newX, tmp, oldX int
	oldX = x
	for x != 0 {
		tmp = x % 10
		newX = newX*10 + tmp
		x = x / 10
	}

	return newX == oldX
}

/*****************************9.回文数end******************************/

//昂，筠，之，齐，扬，溪, 舟，弈, 茗，宁

// RomanToInt /*****************************13. 罗马数字转整数*****************************/
// 例如， 罗马数字 2 写做II，即为两个并列的 1。12 写做XII，即为X+II。 27 写做XXVII, 即为XX+V+II。
// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做IIII，而是IV。数字 1 在数字 5 的左边，
// 所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为IX。这个特殊的规则只适用于以下六种情况：
// I可以放在V(5) 和X(10) 的左边，来表示 4 和 9。
// X可以放在L(50) 和C(100) 的左边，来表示 40 和90。
// C可以放在D(500) 和M(1000) 的左边，来表示400 和900。
// 给定一个罗马数字，将其转换成整数。输入确保在 1到 3999 的范围内。
//
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
//
// 左边代表的值比右边小就减，否则就加，最后一位总是加,貌似和上面的规律没啥关系
// IV : -1+5 = 4
// VI : 5+1 = 6
// IXL : -1-10+50 = 39
// IVX : -1-5+10 = 4
// IXV : -1+10+5 = 14
// IXVV : -1+10+5+5 = 14
// 大的在后 前面的 减
// 小的在后 前面的 加
// 最后一位总是加
func RomanToInt(s string) int {
	tm := make(map[byte]int)
	tm['I'] = 1
	tm['V'] = 5
	tm['X'] = 10
	tm['L'] = 50
	tm['C'] = 100
	tm['D'] = 500
	tm['M'] = 1000

	//['L',"V","I","I","I"]
	//['L',"V","X"]
	var res, i, sLen int
	for i, sLen = 0, len(s)-1; i < sLen; i++ {
		if tm[s[i+1]] > tm[s[i]] {
			res -= tm[s[i]]
		} else {
			res += tm[s[i]]
		}
	}

	// 最后一位不用比较，都是 加
	res += tm[s[i]]

	return res
}

/*****************************13. 罗马数字转整数*****************************/

// LongestCommonPrefix /*****************************14.最长公共前缀*****************************/
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
//
// 输入: ["flower","flow","flight"]
// 输出: "fl"
//
// 输入: ["dog","racecar","car"]
// 输出: ""
// 解释: 输入不存在公共前缀。s
//
// 所有输入只包含小写字母 a-z 。
// 设第一个字符串为公共前缀，然后用公共前缀依次比较之后的字符串，更新公共前缀
// 公共长度一定小于等于最短长度
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
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
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
	if s[0] == ')' || s[0] == ']' || s[0] == '}' || len(s) <= 1 {
		return false
	}

	b := make([]byte, 0) //存储左边括号的字符切片
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			b = append(b, s[i]) //左括号进数组
			continue
		}
		if len(b) == 0 {
			return false
		}
		switch s[i] {
		case ')':
			if b[len(b)-1] != '(' {
				return false
			}
		case ']':
			if b[len(b)-1] != '[' {
				return false
			}
		case '}':
			if b[len(b)-1] != '{' {
				return false
			}
		default:
			return false
		}
		b = b[:len(b)-1] //切片最后一位出队
	}
	return len(b) == 0
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
// 方法1, 定义新链表, 有头尾指针
func MergeTwoListsMy(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	head := &ListNode{}
	tail := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			tail.Next = list2
			return head.Next
		}

		if list2 == nil {
			tail.Next = list1
			return head.Next
		}

		if list1.Val > list2.Val {
			tail.Next = list2
			list2 = list2.Next
			tail = tail.Next
			tail.Next = nil
		} else {
			tail.Next = list1
			list1 = list1.Next
			tail = tail.Next
			tail.Next = nil
		}
	}
	return head
}

// MergeTwoListTwoPoint 方法2：双指针法, 效率更高
func MergeTwoListTwoPoint(list1 *ListNode, list2 *ListNode) *ListNode {
	//初始化一个虚拟的头节点
	newList := &ListNode{}
	p := newList
	p1 := list1
	p2 := list2
	//遍历对比两个指针值的大小，有一个走完了就停止
	for p1 != nil && p2 != nil {
		//将值小的节点接到p指针后面
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		//p指针前进, 但是newList保留在头部,所以 newList 的下一个就是真正的链表头
		p = p.Next
	}
	//将未比较的剩余节点都放到p指针后面
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	//返回虚拟头节点的下一个节点就是真正的头节点
	return newList.Next
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

// RemoveDuplicatesMap map法, 无序的数组也可以
// []int{10, 3, 2, 1, 1, 1, 6, 2, 3, 3, 4, 4, 6}
func RemoveDuplicatesMap(nums []int) int {
	newArr := make([]int, 0)
	m := make(map[int]bool)
	for _, v := range nums {
		_, ok := m[v]
		if !ok {
			newArr = append(newArr, v)
			m[v] = true
		}
	}
	fmt.Println("newArr: ", newArr)
	return len(newArr)
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
