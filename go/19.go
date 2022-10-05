package main

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
// Related Topics 链表 双指针 👍 2247 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	list := &ListNode{0, head}
	temp := list
	for i := 0; i < length-n; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return list.Next
}

//leetcode submit region end(Prohibit modification and deletion)
