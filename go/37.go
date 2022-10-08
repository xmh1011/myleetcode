package main

//编写一个程序，通过填充空格来解决数独问题。
//
// 数独的解法需 遵循如下规则：
//
//
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//
//
// 数独部分空格内已填入了数字，空白格用 '.' 表示。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".
//",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".
//","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6
//"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[
//".",".",".",".","8",".",".","7","9"]]
//输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8
//"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],[
//"4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9",
//"6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4",
//"5","2","8","6","1","7","9"]]
//解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
//
//
//
//
//
//
//
//
//
// 提示：
//
//
// board.length == 9
// board[i].length == 9
// board[i][j] 是一位数字或者 '.'
// 题目数据 保证 输入数独仅有一个解
//
//
// Related Topics 数组 回溯 矩阵 👍 1420 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	rows, columns := [9][10]int{}, [9][10]int{} //构建行和列的数组
	subboxs := [3][3][10]int{}                  // 构建方块的数组
	perhaps := [9][9][10]int{}
	counter := [9][9]int{}
	var flag int
	// 2
	for i, row := range board {
		for j, num := range row {
			// 空格不用管
			if num == '.' {
				continue
			}
			number := num - '1' //将byte换算成数值
			// 没有出现过，则找打对应的行列宫格map中对应的number次数++
			rows[i][number]++
			columns[j][number]++
			subboxs[i/3][j/3][number]++
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// 记录空格中可能会填入的数字
			for board[i][j] == '.' {
				for k := 1; k < 10; k++ {
					if rows[i][k] == 0 && columns[j][k] == 0 && subboxs[i/3][j/3][k] == 0 {
						perhaps[i][j][k] = 1 // 记录空格中可能会填入的数字
						counter[i][j]++      // 记录可填入的数字的可能数
						flag = k
					}
				}
				if counter[i][j] == 1 {
					board[i][j] = byte(flag + '1')
				}
			}
		}
	}
}

// TODO:不会 抄答案的
// 现在应该是懂了
func solveSudoku_leetcode(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for digit := byte(0); digit < 9; digit++ {
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				if dfs(pos + 1) {
					return true
				}
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
}

//leetcode submit region end(Prohibit modification and deletion)
