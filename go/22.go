package main

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 字符串 动态规划 回溯 👍 2895 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//type stack struct {
//	ele []byte
//	top int
//}
//
//func NewStack(s *stack) *stack {
//	s.top = -1
//	s.ele = []byte{}
//	return s
//}
//func InStack(s *stack, char byte) *stack {
//	s.top = s.top + 1
//	s.ele[s.top] = char
//	return s
//}
//func OutStack(s *stack) (*stack, byte) {
//	char := s.ele[s.top]
//	s.top = s.top - 1
//	return s, char
//}
//func generateParenthesis(n int) []string {
//	var left_stack, right_stack *stack
//	left_stack = NewStack(left_stack)
//	right_stack = NewStack(right_stack)
//	for i := 0; i < n; i++ {
//		InStack(left_stack, '(')
//		InStack(right_stack, ')')
//	}
//	ans := []string{}
//
//	return ans
//}

// TODO review dfs
func generateParenthesis(n int) []string {
	var ret []string
	var dfs func(int, int, []byte)
	dfs = func(r, l int, path []byte) { //1)同17.电话号码的字母组合
		if r == 0 && l == 0 {
			ret = append(ret, string(path))
			return
		}
		if r > 0 {
			dfs(r-1, l+1, append(path, '('))
		}
		if l > 0 {
			dfs(r, l-1, append(path, ')'))
		}
	}
	dfs(n, 0, make([]byte, 0, 2*n))
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
