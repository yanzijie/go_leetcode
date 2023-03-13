package code

import "fmt"

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
