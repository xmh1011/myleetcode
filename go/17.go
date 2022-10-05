package main

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
//
// Related Topics 哈希表 字符串 回溯 👍 2136 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type button struct {
	number byte
	letter string
}

var phonebutton = &[]button{
	{'1', ""},
	{'2', "abc"},
	{'3', "def"},
	{'4', "ghi"},
	{'5', "jkl"},
	{'6', "nmo"},
	{'7', "pqrs"},
	{'8', "tuv"},
	{'9', "wxyz"},
}

//func letterCombinations(digits string) []string {
//	len := len(digits)
//	result := []string{}
//	words := []string{}
//	length := []int{}
//	for i := 0; i < len; i++ {
//		length[i], words[i] = processLetter(digits[i])
//	}
//}
//
//func processLetter(letter byte) (length int, s string) {
//	if letter == '1' {
//		s = ""
//	}
//	if letter == '2' {
//		s = "abc"
//	}
//	if letter == '3' {
//		s = "def"
//	}
//	if letter == '4' {
//		s = "ghi"
//	}
//	if letter == '5' {
//		s = "jkl"
//	}
//	if letter == '6' {
//		s = "mno"
//	}
//	if letter == '7' {
//		s = "pqrs"
//	}
//	if letter == '8' {
//		s = "tuv"
//	}
//	if letter == '9' {
//		s = "wxyz"
//	}
//	return len(s), s
//}

//var phoneMap map[string]string = map[string]string{
//	"2": "abc",
//	"3": "def",
//	"4": "ghi",
//	"5": "jkl",
//	"6": "mno",
//	"7": "pqrs",
//	"8": "tuv",
//	"9": "wxyz",
//}
//
//var combinations []string
//
//func letterCombinations(digits string) []string {
//	if len(digits) == 0 {
//		return []string{}
//	}
//	combinations = []string{}
//	backtrack(digits, 0, "")
//	return combinations
//}
//
//func backtrack(digits string, index int, combination string) {
//	if index == len(digits) {
//		combinations = append(combinations, combination)
//	} else {
//		digit := string(digits[index])
//		letters := phoneMap[digit]
//		lettersCount := len(letters)
//		for i := 0; i < lettersCount; i++ {
//			backtrack(digits, index+1, combination+string(letters[i]))
//		}
//	}
//}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var mapping = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var ret []string
	var dfs func(int, []byte)
	dfs = func(p int, path []byte) { //闭包digits,ret。
		if p == len(digits) {
			ret = append(ret, string(path)) //只在最后进行path的序列化，没有额外的空间浪费。
			return
		}
		d := digits[p]
		for _, c := range mapping[d] {
			dfs(p+1, append(path, c)) //充分利用append的特性，不会导致空间浪费
		}
	}
	dfs(0, make([]byte, 0, len(digits)))
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
