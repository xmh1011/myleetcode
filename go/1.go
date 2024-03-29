package main

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 你可以按任意顺序返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
//p
// 示例 2：
//
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//
//
// 示例 3：
//
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
// 只会存在一个有效答案
//
//
// 进阶：你可以想出一个时间复杂度小于 O(n²) 的算法吗？
//
// Related Topics 数组 哈希表 👍 15473 👎 0
// TODO hashmap
func twoSum(nums []int, target int) []int {
	hashmap := map[int]int{}
	for i, x := range nums {
		if p, ok := hashmap[target-x]; ok {
			return []int{p, i}
		}
		hashmap[x] = i
	}
	return nil
}

//// write a quick sort function
//func quickSort(nums []int) []int {
//	if len(nums) <= 1 {
//		return nums
//	}
//	pivot := nums[0]
//	var left, right []int
//	for _, x := range nums[1:] {
//		if x < pivot {
//			left = append(left, x)
//		} else {
//			right = append(right, x)
//		}
//	}
//	return append(append(quickSort(left), pivot), quickSort(right)...)
//}
//
//// hash sort function
//func hashSort(nums []int) []int {
//	hashmap := map[int]int{}
//	for _, x := range nums {
//		hashmap[x]++
//	}
//	var res []int
//	for i := 0; i < 10000; i++ {
//		if hashmap[i] > 0 {
//			for j := 0; j < hashmap[i]; j++ {
//				res = append(res, i)
//			}
//		}
//	}
//	return res
//}
