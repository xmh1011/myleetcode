use std::collections::HashSet;

struct Solution;

impl Solution {
    pub fn length_of_longest_substring(s: String) -> usize {
        let mut occ = HashSet::new(); // 使用哈希集合来记录字符是否出现
        let n = s.len(); // 获取字符串长度
        let s_chars: Vec<char> = s.chars().collect(); // 将字符串转换为字符向量
        let mut rk = -1; // 右指针初始化为 -1
        let mut ans = 0; // 存储最长无重复子串的长度

        // 枚举左指针的位置
        for i in 0..n {
            if i > 0 {
                // 左指针向右移动一格，移除一个字符
                occ.remove(&s_chars[i - 1]);
            }

            // 移动右指针，直到遇到重复字符
            while (rk + 1) < n as isize && !occ.contains(&s_chars[(rk + 1) as usize]) {
                rk += 1;
                occ.insert(s_chars[rk as usize]);
            }

            // 计算当前无重复子串的长度
            ans = ans.max((rk - i as isize + 1) as usize);
        }
        ans
    }
}

fn main() {
    // 创建 Solution 对象
    let solution = Solution;

    // 测试用例
    let test_str = "abcabcbb".to_string();
    println!("Input String: {}", test_str);

    // 调用函数并输出结果
    let result = solution.length_of_longest_substring(test_str);
    println!("Length of Longest Substring Without Repeating Characters: {}", result);
}
