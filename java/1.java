import java.util.HashMap;
import java.util.Map;

public class Solution {
    // 实现 two_sum 方法，返回两个数的索引
    public int[] twoSum(int[] nums, int target) {
        // 使用 HashMap 存储数组元素和对应的索引
        Map<Integer, Integer> map = new HashMap<>(nums.length);

        // 遍历数组
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            // 检查是否存在目标值所需的补数
            if (map.containsKey(complement)) {
                return new int[]{map.get(complement), i};
            }
            // 将当前元素加入 HashMap 中
            map.put(nums[i], i);
        }
        // 如果没有找到匹配的两数之和，抛出异常
        throw new IllegalArgumentException("No two sum solution found");
    }

    public static void main(String[] args) {
        Solution solution = new Solution();

        // 测试用例
        int[] nums = {2, 7, 11, 15};  // 输入数组
        int target = 9;               // 目标值

        // 调用 twoSum 方法
        int[] result = solution.twoSum(nums, target);

        // 打印结果
        System.out.printf("Indices: [%d, %d]\n", result[0], result[1]);
    }
}
