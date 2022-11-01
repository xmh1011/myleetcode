package main

//整数数组 nums 按升序排列，数组中的值 互不相同 。
//
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[
//k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2
//,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//
//
// 示例 2：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//
// 示例 3：
//
//
//输入：nums = [1], target = 0
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// nums 中的每个值都 独一无二
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 二分查找 👍 2346 👎 0

// 尝试二分查找
//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	len := len(nums)
	result := findx(0, len-1, target, nums)
	return result
}
func findx(head int, tail int, target int, nums []int) int {
	var middle int
	if (head+tail)%2 == 0 {
		middle = (head + tail) / 2
	} else {
		middle = (head+tail)/2 + 1
	}
	if head != tail {
		if nums[head] < nums[tail] && target > nums[head] && target < nums[tail] {
			if target > nums[middle] {
				return findx(middle, tail, target, nums)
			} else {
				return findx(head, middle-1, target, nums)
			}
		} else if nums[head] < nums[tail] && target > nums[tail] {
			return -1
		} else if nums[head] < nums[tail] && target < nums[head] {
			return -1
		} else if nums[head] > nums[tail] && target > nums[head] {
			if target > nums[middle] {
				return findx(head, middle-1, target, nums)
			} else if nums[middle] > nums[head] && target < nums[middle] {
				return findx(head, middle-1, target, nums)
			} else if nums[tail] < nums[middle] && target > nums[middle] {
				return findx(middle, tail, target, nums)
			} else if nums[head] < nums[middle] && target > nums[middle] {
				return findx(middle, tail, target, nums)
			} else if nums[head] < nums[middle] && target > nums[middle] {
				return findx(middle, tail, target, nums)
			} else if nums[middle] == target {
				return middle
			} else {
				return -1
			}
		} else if nums[head] > nums[tail] && target < nums[tail] {
			if target > nums[middle] {
				return findx(middle, tail, target, nums)
			} else if nums[middle] < nums[tail] && target < nums[middle] {
				return findx(head, middle-1, target, nums)
			} else if nums[tail] < nums[middle] {
				return findx(middle, tail, target, nums)
			} else if nums[head] < nums[middle] {
				return findx(middle, tail, target, nums)
			} else if nums[middle] == target {
				return middle
			} else {
				return -1
			}
		} else if nums[head] > nums[tail] && target > nums[tail] && target < nums[head] {
			return -1
		} else if nums[head] == target {
			return head
		} else if nums[tail] == target {
			return tail
		} else {
			return -1
		}
	} else {
		if target == nums[tail] {
			return tail
		} else {
			return -1
		}
	}
}

func search_leetcode(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
