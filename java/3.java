import java.util.HashSet;
import java.util.Set;

public class Solution {
    public int lengthOfLongestSubstring(String s) {
        // 哈希集合，记录每个字符是否出现过
        Set<Character> occ = new HashSet<>();
        int n = s.length();
        // 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧
        int rk = -1, ans = 0;

        // 枚举左指针的位置
        for (int i = 0; i < n; ++i) {
            if (i != 0) {
                // 左指针向右移动一格，移除一个字符
                occ.remove(s.charAt(i - 1));
            }
            // 不断地移动右指针，直到出现重复字符
            while (rk + 1 < n && !occ.contains(s.charAt(rk + 1))) {
                occ.add(s.charAt(rk + 1));
                ++rk;
            }
            // 第 i 到 rk 个字符是一个极长的无重复字符子串
            ans = Math.max(ans, rk - i + 1);
        }
        return ans;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();

        // 测试用例
        String testStr = "abcabcbb";
        System.out.println("Input String: " + testStr);

        // 调用函数并输出结果
        int result = solution.lengthOfLongestSubstring(testStr);
        System.out.println("Length of Longest Substring Without Repeating Characters: " + result);
    }
}
