package main

import (
	"sort"
)

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
// Related Topics 数组 回溯 👍 1214 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func permuteUnique(nums []int) [][]int {
//	ans = [][]int{}
//	first := []int{nums[0]}
//	ans = [][]int{first}
//	var history map[int]bool
//	history = make(map[int]bool)
//	if len(nums) == 0 {
//		return [][]int{}
//	} else if len(nums) == 1 {
//		return ans
//	} else {
//		for i := 1; i < len(nums); i++ {
//			ans = backTrackUnique(nums[i], ans, i+1, history)
//		}
//	}
//	return ans
//}
//
//func backTrackUnique(num int, ans [][]int, n int, history map[int]bool) [][]int {
//	for i := 0; i < len(ans); i++ {
//		for j := 0; j < n-1; j++ {
//			cur := append(ans[i][:j+1], num)
//			cur = append(cur, ans[i][j+1:]...)
//			ans = append(append(ans[:i], cur), ans[i+1:]...)
//		}
//		ans = append(append(ans[:i], append(ans[i][0:], num)), ans[i+1:]...)
//	}
//	return ans
//}
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
