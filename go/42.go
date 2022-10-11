package main

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 3862 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	sum := 0
	max := getMax(height) //找到最大的高度，以便遍历。
	for i := 1; i <= max; i++ {
		isStart := false //标记是否开始更新 temp
		temp_sum := 0
		for j := 0; j < len(height); j++ {
			if isStart && height[j] < i {
				temp_sum++
			}
			if height[j] >= i {
				sum = sum + temp_sum
				temp_sum = 0
				isStart = true
			}
		}
	}
	return sum
}
func getMax(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
