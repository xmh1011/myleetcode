#include <stdio.h>
#include <stdlib.h>

// 函数声明
int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    for (int i = 0; i < numsSize; ++i) {
        for (int j = i + 1; j < numsSize; ++j) {
            if (nums[i] + nums[j] == target) {
                int* ret = (int*)malloc(sizeof(int) * 2);  // 分配内存
                if (ret == NULL) {  // 检查内存分配是否成功
                    *returnSize = 0;
                    return NULL;
                }
                ret[0] = i;
                ret[1] = j;
                *returnSize = 2;
                return ret;
            }
        }
    }
    *returnSize = 0;
    return NULL;
}

int main() {
    int nums[] = {2, 7, 11, 15};  // 示例数组
    int numsSize = sizeof(nums) / sizeof(nums[0]);  // 计算数组大小
    int target = 9;  // 目标值
    int returnSize;  // 存储返回数组的大小

    // 调用 twoSum 函数
    int* result = twoSum(nums, numsSize, target, &returnSize);

    if (result != NULL) {
        printf("Indices: %d, %d\n", result[0], result[1]);  // 输出结果
        free(result);  // 释放分配的内存
    } else {
        printf("No two sum solution found.\n");
    }

    return 0;
}
