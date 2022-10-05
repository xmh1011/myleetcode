package main

import (
	"sort"
)

//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
//
//
// 你可以按 任意顺序 返回答案 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
//
//
// Related Topics 数组 双指针 排序 👍 1392 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 尝试固定两个指针，再用两个移动指针
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	if len(nums) < 4 {
		return ans
	}
	for i := 0; i < len(nums)-3; {
		for j := i + 1; j < len(nums)-2; {
			left := j + 1
			right := len(nums) - 1
			for left < right {
				if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					left = sort.SearchInts(nums, nums[left]+1)
				}
			}
			j = sort.SearchInts(nums, nums[j]+1)
		}
		i = sort.SearchInts(nums, nums[i]+1)
	}
	return ans
}

//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
//
//
// 你可以按 任意顺序 返回答案 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
//
//
// Related Topics 数组 双指针 排序 👍 1392 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//// 尝试固定两个指针，再用两个移动指针
//func fourSum(nums []int, target int) [][]int {
//	sort.Ints(nums)
//	ans := make([][]int, 0)
//	if len(nums) < 4 {
//		return ans
//	}
//	for first := 0; first < len(nums)-3; first++ {
//		for forth := len(nums) - 1; forth > first+2; forth-- {
//			third := forth - 1
//			second := first + 1
//			if nums[second]+nums[third] < target-nums[forth]-nums[first] && second < third {
//				second++
//			} else if nums[second]+nums[third] > target-nums[forth]-nums[first] && second < third {
//				third--
//			} else if nums[second]+nums[third] == target-nums[forth]-nums[first] && second < third {
//				ans = append(ans, []int{nums[first], nums[second], nums[third], nums[forth]})
//				//if nums[second] != nums[second+1] && second < third-1 {
//				//	second++
//				//}
//				//if nums[third] != nums[third-1] && second < third-1 {
//				//	third--
//				//}
//			}
//		}
//	}
//	return ans
//}

//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region end(Prohibit modification and deletion)
