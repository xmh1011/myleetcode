package main

//给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
//
//
// 示例 2：
//
//
//输入：nums = [3,2,1,0,4]
//输出：false
//解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// 0 <= nums[i] <= 10⁵
//
//
// Related Topics 贪心 数组 动态规划 👍 2086 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[0] == 0 && len(nums) != 1 {
			return false
		} else if nums[i] == 0 && i != 0 && i != len(nums)-1 {
			for j := i - 1; j >= 0; j-- {
				if nums[j] > i-j {
					break
				} else if nums[0] <= i && j == 0 {
					return false
				}
			}
		}
	}
	return true
}

func canJump_GreedyAlgorithm(nums []int) bool {
	rightmost := 0
	for i := 0; i < len(nums); i++ {
		if i <= rightmost {
			rightmost = max(rightmost, i+nums[i])
			if rightmost >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
