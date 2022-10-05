package main

//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
//
//提示：
//
//
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
//
//
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
//
//
//
//
// Related Topics 递归 链表 👍 1821 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	cur := ListNode{0, head}
//	for i := 0; i < getlength(head)/k; i++ {
//		first := ListNode{0, head}
//		head = head.Next
//		for n := 0; n < k-1; n++ {
//			head.Next = first.Next
//			first.Next = head
//			head = head.Next
//		}
//	}
//	return cur.Next
//}
//
////func reverseListNode(head *ListNode, k int) {
////	first := ListNode{0, head}
////	head = head.Next
////	for i := 0; i < k-1; i++ {
////		head.Next = first.Next
////		first.Next = head
////		head = head.Next
////	}
////	return
////}
//
//func getlength(head *ListNode) int {
//	len := 0
//	for head.Next != nil {
//		len++
//		head = head.Next
//	}
//	return len
//}
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

//leetcode submit region end(Prohibit modification and deletion)
