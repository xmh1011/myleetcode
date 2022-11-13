package main

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
//
// Related Topics 数组 矩阵 模拟 👍 1259 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func spiralOrder(matrix [][]int) []int {
//	ans := []int{}
//	rows, columns := len(matrix), len(matrix[0])
//	ans = make([]int, rows*columns)
//	if rows == 0 || columns == 0 {
//		return []int{}
//	} else if rows == 1 {
//		for i := 0; i < columns; i++ {
//			ans = append(ans, matrix[0][i])
//		}
//		return ans
//	} else if columns == 1 {
//		for i := 0; i < rows; i++ {
//			ans = append(ans, matrix[i][0])
//		}
//		return ans
//	} else {
//		for j := 0; j < rows; j = j + 2 {
//			for i := 0; i < columns; i++ {
//				ans = append(ans, matrix[0][i])
//			}
//			for i := 1; i < rows; i++ {
//				ans = append(ans, matrix[i][columns-1])
//			}
//			for i := columns - 2; i >= 0; i-- {
//				ans = append(ans, matrix[rows-1][i])
//			}
//			for i := rows - 2; i > 0; i-- {
//				ans = append(ans, matrix[i][0])
//			}
//			ans = append(ans, spiralOrder(matrix[j/2 : rows-j/2][j/2:columns-j/2])...)
//		}
//		return ans
//	}
//}
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

//leetcode submit region end(Prohibit modification and deletion)
