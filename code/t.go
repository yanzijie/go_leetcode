package code

import (
	"fmt"
	"sort"
)

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
	// 这个算法效率最高，但是没有修改原数组
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

func RemoveDuplicatesSlice(nums []int) int {
	// 切片自己删自己
	var i int
	for i = 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i:]...)
			i--
		}
	}
	return len(nums)
}

/******************************26. 删除有序数组中的重复项******************************/

// RemoveElement ******************************27. 移除元素******************************
/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度
例如:
输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3]
解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。
*/
func RemoveElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

/******************************27. 移除元素******************************/

// SearchInsert /******************************35. 搜索插入位置******************************/
/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。

输入: nums = [1,3,5,6], target = 5
输出: 2

输入: nums = [1,3,5,6], target = 2
输出: 1

输入: nums = [1,3,5,6], target = 7
输出: 4
*/
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

/******************************35. 搜索插入位置******************************/

// FindDifferentElements /*********************************非leetcode题目************************************/
/*
实现一个"找出两个数组中不同元素"的算法，需要考虑两个数组元素数量不同的情况
例如:  输入[1,2,3,4,5,6,7] 和 [1,2,3,5,6], 返回 [4,7]
*/
/*********************************非leetcode题目************************************/
func FindDifferentElements(arr1 []int, arr2 []int) []int {
	res := make([]int, 0)

	// 把arr1数组中不在arr2中的元素添加到res中
	for _, item := range arr1 {
		if !isInArr(arr2, item) {
			res = append(res, item)
		}
	}
	// 把arr2数组中不在arr1中的元素添加到res中
	for _, item := range arr2 {
		if !isInArr(arr1, item) {
			res = append(res, item)
		}
	}

	return res
}

// 判断元素 item 是否在数组 arr 中, 在返回true, 不在false
func isInArr(arr []int, item int) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

// RemoveDuplicatesString /*********************************非leetcode题目************************************/
/*
数组元素去重
*/
func RemoveDuplicatesString(strs []string) []string {
	// 如果数组长度小于等于1，则无需进行去重操作，直接返回原始数组
	if len(strs) <= 1 {
		return strs
	}

	// 先将数组排序，方便后面去重操作
	sort.Strings(strs)

	// 定义慢指针i和快指针j，初始值分别为0和1
	i := 0
	for j := 1; j < len(strs); j++ {
		// 不相等，则将nums[j]赋值给nums[i+1],覆盖重复的元素，并将i向右移动一位
		if strs[j] != strs[i] {
			// 由于 i++ 了，所以在正常情况下 i 和 j 是相等的，所以覆盖不会产生问题
			// 只有要有一次前后元素都相等，i就不会向右移动,
			// i和j这个时候才不相等，这个时候就可以覆盖了
			i++
			strs[i] = strs[j]
		}
	}

	// 返回删除后的数组，数组长度为i+1, i之后的元素就不要了
	return strs[0 : i+1]
}

func RemoveDuplicatesInt(nums []int) []int {
	// 如果数组长度小于等于1，则无需进行去重操作，直接返回原始数组
	if len(nums) <= 1 {
		return nums
	}

	// 先将数组排序，方便后面去重操作
	sort.Ints(nums)

	// 定义慢指针i和快指针j，初始值分别为0和1
	i := 0
	for j := 1; j < len(nums); j++ {
		// 不相等，则将nums[j]赋值给nums[i+1],覆盖重复的元素，并将i向右移动一位
		if nums[j] != nums[i] {
			// 由于 i++ 了，所以在正常情况下 i 和 j 是相等的，所以覆盖不会产生问题
			// 只有要有一次前后元素都相等，i就不会向右移动,
			// i和j这个时候才不相等，这个时候就可以覆盖了
			i++
			nums[i] = nums[j]
		}
	}

	// 返回删除后的数组，数组长度为i+1, i之后的元素就不要了
	return nums[0 : i+1]
}

/*********************************非leetcode题目************************************/

// PlusOne /******************************66. 加一******************************/
/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

输入: digits = [1,2,3]  输入数组表示数字 123。
输出: [1,2,4]

输入: digits = [4,3,2,1]  输入数组表示数字 4321。
输出: [4,3,2,2]

输入: digits = [0]
输出: [1]

1 <= digits.length <= 100
0 <= digits[i] <= 9
*/
func PlusOne(digits []int) []int {
	carry := 1 // 表示进位，初始为 1，因为要在原数字的基础上加 1

	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10 // 需要进位的时候,sum=10, 10%10=1, 0-9的数 % 10都等于自身
		carry = sum / 10     // 如果sum等于10，这个位等于0，那么进位还得保留是1, 999->1000
	}

	// 如果最高位有进位，需要在数组头部插入进位值
	if carry > 0 {
		// 到了这里说明是新数组长度超过了原来的长度，例如 [9,9,9]->[1,0,0,0]
		// append( [1], [0,0,0] ) = [1,0,0,0]
		digits = append([]int{carry}, digits...)
	}

	return digits
}

/******************************35. 搜索插入位置******************************/
