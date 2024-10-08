use std::collections::HashMap;

// 定义一个 `Solution` 结构体，用于实现 `two_sum` 函数
struct Solution;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map = HashMap::with_capacity(nums.len());

        // 遍历数组并查找目标值
        for (i, &num) in nums.iter().enumerate() {
            if let Some(&k) = map.get(&(target - num)) {
                return vec![k as i32, i as i32];
            }
            map.insert(num, i);
        }
        // 如果没有找到匹配的两数之和，返回一个空的向量
        panic!("No two sum solution found");
    }
}

fn main() {
    // 测试用例
    let nums = vec![2, 7, 11, 15];  // 输入数组
    let target = 9;                 // 目标值

    // 调用 `two_sum` 方法
    let result = Solution::two_sum(nums, target);

    // 打印结果
    println!("Indices: {:?}", result);
}
