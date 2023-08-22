package code

import (
	"strconv"
	"strings"
)

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

// AddBinary *****************************67. 二进制求和*****************************
/*
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

输入:a = "11", b = "1"
输出："100"

输入：a = "1010", b = "1011"
输出："10101"
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

//LengthOfLastWord /******************************58. 最后一个单词的长度Start******************************/
/*
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。

输入：s = "Hello World"
输出：5
解释：最后一个单词是“World”，长度为5。

输入：s = "   fly me   to   the moon  "
输出：4
解释：最后一个单词是“moon”，长度为4。
*/
func LengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	return len(arr[len(arr)-1])

	// 不用 strings.TrimSpace(s) 的方案
	//na := arr[:0]
	//for i := 0; i < len(arr); i++ {
	//	if len(arr[i]) != 0 {
	//		na = append(na, arr[i])
	//	}
	//}
	//return len(na[len(na)-1])
}

/******************************58. 最后一个单词的长度End******************************/

//StrStr /******************************28. 找出字符串中第一个匹配项的下标 Start******************************/
/*
给你两个字符串 haystack 和 needle ，
请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）,
如果 needle 不是 haystack 的一部分，则返回  -1 。

示例 1：
输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0

示例 2：
输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1
*/
func StrStr(haystack string, needle string) int {
	// 使用系统包...
	//return strings.Index(haystack, needle)

	if len(haystack) == 0 || len(needle) == 0 || len(needle) > len(haystack) {
		return -1
	}

	// 因为存在needle只是一个字符的情况，所以 i<=len 最后一个也要看
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// 从第一个元素开始，以len(needle)为长度, 一组一组的对比
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

/******************************28. 找出字符串中第一个匹配项的下标 End******************************/
