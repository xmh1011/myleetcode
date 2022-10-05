package main

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//
// 例如，121 是回文，而 123 不是。
//
//
//
//
// 示例 1：
//
//
//输入：x = 121
//输出：true
//
//
// 示例 2：
//
//
//输入：x = -121
//输出：false
//解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
//
// 示例 3：
//
//
//输入：x = 10
//输出：false
//解释：从右向左读, 为 01 。因此它不是一个回文数。
//
//
//
//
// 提示：
//
//
// -2³¹ <= x <= 2³¹ - 1
//
//
//
//
// 进阶：你能不将整数转为字符串来解决这个问题吗？
//
// Related Topics 数学 👍 2242 👎 0

// 这么笨的算法一看就是我自己写的
func isPalindrome(x int) bool {
	var length, i int
	var num = [10]int{0}
	length = 0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	} else if x < 10 {
		return true
	} else {
		for i = 0; x >= 10; i++ {
			num[i] = x % 10
			x = x / 10
			length++
			if x < 10 {
				num[i+1] = x
			}
		}
		length++
		for i = 0; i <= length/2; i++ {
			if num[i] != num[length-i-1] {
				return false
			}
		}
		return true
	}
}
