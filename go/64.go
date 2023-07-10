package main

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
//
//
// 示例 1：
//
//
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
//
//
// 示例 2：
//
//
// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 200
//
//
// Related Topics 数组 动态规划 矩阵 👍 1530 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	column, row := len(grid[0]), len(grid)
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, column)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			} else if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				continue
			}
		}
	}
	return dp[row-1][column-1]
}

func minPathSumLeetCode(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, columns)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][columns-1]
}

// leetcode submit region end(Prohibit modification and deletion)
