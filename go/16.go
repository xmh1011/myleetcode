package main

import (
	"sort"
)

//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
// 示例 2：
//
//
//输入：nums = [0,0,0], target = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 双指针 排序 👍 1258 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 这题是真简单但是也真恶心，考验认真程度了属于是
func threeSumClosest(nums []int, target int) int {
	result := 10000
	distance := 10000
	sort.Ints(nums)
	len := len(nums)
	first := 0
	for first = 0; first < len-2; first++ {
		third := len - 1
		second := first + 1
		for third > second {
			sum := nums[second] + nums[first] + nums[third]
			if sum < target {
				if distance >= target-sum {
					distance = target - sum
					result = nums[first] + nums[second] + nums[third]
				}
				second++
			}
			if sum >= target {
				if distance >= sum-target {
					distance = sum - target
					result = nums[first] + nums[second] + nums[third]
				}
				third--
			}
		}
	}
	return result
}

func absoluteValue(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

//leetcode submit region end(Prohibit modification and deletion)
