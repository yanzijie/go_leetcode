package code

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
