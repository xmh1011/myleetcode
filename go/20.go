package main

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
//
// Related Topics 栈 字符串 👍 3533 👎 0

// 老早之前写的算法，那时候还很稚嫩，写的不好看
func isValid(s string) bool {
	stack := [10000]byte{}
	len := len(s)
	var flag int
	i := 0
	if len%2 == 1 {
		return false
	} else {
		for _, x := range s {
			if x == '(' || x == '{' || x == '[' {
				stack[i] = byte(x)
				i++
			} else {
				if i == 0 {
					flag = 0
					break
				} else {
					flag, i = match(stack[i-1], byte(x), i)
					if flag == 0 {
						break
					}
				}
			}
		}
		if i != 0 {
			flag = 0
		}
	}
	if flag == 1 {
		return true
	} else {
		return false
	}
}
func match(a, b byte, i int) (int, int) {
	if a == '{' && b == '}' {
		i--
		return 1, i
	} else if a == '(' && b == ')' {
		i--
		return 1, i
	} else if a == '[' && b == ']' {
		i--
		return 1, i
	} else {
		i--
		return 0, i
	}
}
