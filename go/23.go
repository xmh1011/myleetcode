package main

//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2195 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// TODO: implement
func mergeKLists(lists []*ListNode) *ListNode {
	nums := len(lists)
	var ans *ListNode
	if nums == 0 {
		return nil
	} else if nums == 1 {
		return lists[0]
	} else {
		for i := 1; i < nums; i++ {
			ans = mergeLists(lists[0], lists[i])
			lists[0] = ans
		}
		return ans
	}
}

func mergeLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeLists(list1, list2.Next)
		return list2
	}
}

// 下面为leetcode解法
//func mergeKLists(lists []*ListNode) *ListNode {
//	return merge(lists, 0, len(lists)-1)
//}
//
//func merge(lists []*ListNode, l, r int) *ListNode{
//	if l == r {
//		return lists[l]
//	}
//	if l > r {
//		return nil
//	}
//	mid := (l + r) / 2
//	return mergeTwoLists(merge(lists, l , mid), merge(lists, mid + 1, r))
//}
//
//func mergeTwoLists(a, b *ListNode)*ListNode{
//	if a == nil && b == nil{
//		return nil
//	}
//	var head = &ListNode{}
//	dummy := head
//	for a != nil && b != nil{
//		if a.Val < b.Val{
//			head.Next = a
//			a = a.Next
//		}else{
//			head.Next = b
//			b = b.Next
//		}
//		head = head.Next
//	}
//	if a != nil{
//		head.Next = a
//	}else if b != nil{
//		head.Next = b
//	}
//	return dummy.Next
//}

//leetcode submit region end(Prohibit modification and deletion)
