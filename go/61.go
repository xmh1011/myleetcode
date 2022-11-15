package main

//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[4,5,1,2,3]
//
//
// 示例 2：
//
//
//输入：head = [0,1,2], k = 4
//输出：[2,0,1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 10⁹
//
//
// Related Topics 链表 双指针 👍 853 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	} // 考虑链表为空的情况
	n := 1
	list := head
	for list.Next != nil {
		list = list.Next
		n++
	} // 求链表长度
	add := n - k%n
	if add == n {
		return head
	}
	list.Next = head
	for add > 0 {
		list = list.Next
		add--
	}
	ret := list.Next
	list.Next = nil
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
