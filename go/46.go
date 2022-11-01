package main

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
//
// Related Topics 数组 回溯 👍 2251 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func permute_leetcode(nums []int) [][]int {
	ans = [][]int{}
	backTrack_leetcode(nums, len(nums), []int{})
	return ans
}
func backTrack_leetcode(nums []int, numsLen int, path []int) {
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		// The copy built-in function copies elements from a source slice into a
		// destination slice.
		// func copy(dst, src []Type) int
		ans = append(ans, p)
	}
	for i := 0; i < numsLen; i++ {
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...) //直接使用切片
		backTrack_leetcode(nums, len(nums), path)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) //回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]
	}
}

var ans [][]int

func permute(nums []int) [][]int {
	ans = [][]int{}
	first := []int{nums[0]}
	ans = [][]int{first}
	if len(nums) == 0 {
		return [][]int{}
	} else if len(nums) == 1 {
		return ans
	} else {
		for i := 1; i < len(nums); i++ {
			ans = backTrack(nums[i], ans, i+1)
		}
	}
	return ans
}

func backTrack(num int, ans [][]int, n int) [][]int {
	for i := 0; i < len(ans); i++ {
		for j := 0; j < n-1; j++ {
			cur := append(ans[i][:j+1], num)
			cur = append(cur, ans[i][j+1:]...)
			ans = append(append(ans[:i], cur), ans[i+1:]...)
		}
		ans = append(append(ans[:i], append(ans[i][0:], num)), ans[i+1:]...)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
