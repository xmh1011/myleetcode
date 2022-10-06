package main

import (
	"math"
)

//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
// 返回被除数 dividend 除以除数 divisor 得到的商。
//
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//
//
//
// 示例 1:
//
// 输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
//
// 示例 2:
//
// 输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = truncate(-2.33333..) = -2
//
//
//
// 提示：
//
//
// 被除数和除数均为 32 位有符号整数。
// 除数不为 0。
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2³¹, 231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
//
//
// Related Topics 位运算 数学 👍 991 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func divide(dividend int, divisor int) int {
//	var i int
//	if dividend >= 0 && divisor < 0 {
//		return -divide(dividend, -divisor)
//	} else if dividend < 0 && divisor < 0 {
//		return divide(-dividend, -divisor)
//	} else if dividend < 0 && divisor > 0 {
//		return -divide(-dividend, divisor)
//	} else {
//		for i = 0; dividend-divisor >= 0; i++ {
//			dividend = dividend - divisor
//		}
//		if i > math.MaxInt32/10 {
//			return math.MaxInt32 / 10
//		} else {
//			return i
//		}
//	}
//}
// 快速乘
// x 和 y 是负数，z 是正数
// 判断 z * y >= x 是否成立
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 { // 不能使用除法
		if z&1 > 0 {
			// 需要保证 result + add >= x
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 {
			// 需要保证 add + add >= x
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

func divide(dividend, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1 // 注意溢出，并且不能使用除法
		if quickAdd(divisor, mid, dividend) {
			ans = mid
			if mid == math.MaxInt32 { // 注意溢出
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
