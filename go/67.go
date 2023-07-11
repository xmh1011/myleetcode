package main

import (
	"strconv"
)

// 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
//
//
//
// 示例 1：
//
//
// 输入:a = "11", b = "1"
// 输出："100"
//
// 示例 2：
//
//
// 输入：a = "1010", b = "1011"
// 输出："10101"
//
//
//
// 提示：
//
//
// 1 <= a.length, b.length <= 10⁴
// a 和 b 仅由字符 '0' 或 '1' 组成
// 字符串如果不是 "0" ，就不含前导零
//
//
// Related Topics 位运算 数学 字符串 模拟 👍 1049 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)
	
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// leetcode submit region end(Prohibit modification and deletion)
