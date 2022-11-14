package main

import (
	"sort"
)

//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返
//回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2：
//
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//
//
// 提示：
//
//
// 1 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= starti <= endi <= 10⁴
//
//
// Related Topics 数组 排序 👍 1707 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//func merge(intervals [][]int) [][]int {
//	head := make([]int, len(intervals))
//	tail := make([]int, len(intervals))
//	for i := 0; i < len(intervals); i++ {
//		head[i] = intervals[i][0]
//		tail[i] = intervals[i][1]
//	}
//	sort.Ints(head)
//	sort.Ints(tail)
//	var ans [][]int
//	for i := 0; i < len(intervals); i++ {
//		for j := i; j < len(intervals); j++ {
//			if tail[i] > head[j] && tail[i] < head[j+1] { //得考虑相等的情况
//				ans = append(ans,{head[j],tail[i]})
//				break
//			}
//		}
//	}
//}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 没有一点重合
			res = append(res, prev)
			prev = cur
		} else { // 有重合
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
