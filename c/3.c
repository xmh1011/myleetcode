#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

// 查找最长无重复字符子串的长度
int lengthOfLongestSubstring(char* s) {
    // 使用布尔数组记录每个字符是否出现过
    bool occ[256] = {false};  // 假设输入为标准 ASCII 码，共 256 个字符
    int n = strlen(s);  // 计算字符串长度
    int rk = -1;  // 右指针，初始值为 -1
    int ans = 0;  // 最长无重复子串的长度

    // 枚举左指针的位置
    for (int i = 0; i < n; ++i) {
        if (i != 0) {
            // 左指针向右移动一格，移除一个字符
            occ[(unsigned char)s[i - 1]] = false;
        }
        // 不断地移动右指针，直到遇到重复字符为止
        while (rk + 1 < n && !occ[(unsigned char)s[rk + 1]]) {
            occ[(unsigned char)s[rk + 1]] = true;
            ++rk;
        }
        // 第 i 到 rk 个字符是一个极长的无重复字符子串
        if (ans < rk - i + 1) {
            ans = rk - i + 1;
        }
    }
    return ans;
}

int main() {
    // 测试用例
    char testStr[] = "abcabcbb";
    printf("Input String: %s\n", testStr);

    // 调用函数并输出结果
    int result = lengthOfLongestSubstring(testStr);
    printf("Length of Longest Substring Without Repeating Characters: %d\n", result);

    return 0;
}
