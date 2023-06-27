package main

import (
	"math/big"
)

// 一个机器人位于一个
// m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
//
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//
//
// 示例 1：
//
//
// 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// 输出：2
// 解释：3x3 网格的正中间有一个障碍物。
// 从左上角到右下角一共有 2 条不同的路径：
// 1. 向右 -> 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右 -> 向右
//
//
// 示例 2：
//
//
// 输入：obstacleGrid = [[0,1],[0,0]]
// 输出：1
//
//
//
//
// 提示：
//
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] 为 0 或 1
//
//
// Related Topics 数组 动态规划 矩阵 👍 1071 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	
	// 初始化
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	// 边界条件
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}
	
	// 状态转移
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 0 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	
	// 返回结果
	return dp[m-1][n-1]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	var x, y int
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				x = i
				y = j
			}
		}
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	path := int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())                   // C(m+n-2, n-1)
	pathToObstacle := int(new(big.Int).Binomial(int64(x+y), int64(y)).Int64())             // C(i+j, j)
	pathFromObstacle := int(new(big.Int).Binomial(int64(m+n-x-y-2), int64(n-y-1)).Int64()) // C(m+n-i-j-2, n-j-1)
	return path - pathToObstacle*pathFromObstacle
}

// leetcode submit region end(Prohibit modification and deletion)
