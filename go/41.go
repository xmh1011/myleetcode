package main

import (
	"fmt"
)

//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。 请你实现时间复杂度为
//O(n) 并且只使用常数级别额外空间的解决方案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,0]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [3,4,-1,1]
//输出：2
//
//
// 示例 3：
//
//
//输入：nums = [7,8,9,11,12]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
// Related Topics 数组 哈希表 👍 1629 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func firstMissingPositive(nums []int) int {
//	sort.Ints(nums)
//	var ans int
//	flag := false
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i+1]-nums[i] > 1 {
//			if nums[i+1] > 0 && nums[i] < 0 && nums[i+1] != 1 {
//				ans = 1
//				flag = true
//				break
//			} else if nums[i+1] > 0 && nums[i] > 0 {
//				ans = nums[i] + 1
//				flag = true
//				break
//			} else if nums[i+1] < 0 && nums[i] < 0 {
//				continue
//			} else if nums[i+1] > 0 && nums[i] < 0 && nums[i+1] == 1 {
//				continue
//			}
//
//		}
//	}
//	if nums[0] > 1 || nums[len(nums)-1] < 1 {
//		return 1
//	} else if flag {
//		return ans
//	} else {
//		return nums[len(nums)-1] + 1
//	}
//}
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			fmt.Println(num - 1)
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 其实这个方法是最好的，眼前一亮那种
func firstMissingPositive_leetcode(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

//leetcode submit region end(Prohibit modification and deletion)
