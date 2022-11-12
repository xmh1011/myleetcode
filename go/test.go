package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := ListNode{0, head}
	for i := 0; i < getlength(head)/k; i++ {
		first := ListNode{0, head}
		head = head.Next
		for n := 0; n < k-1; n++ {
			head.Next = first.Next
			first.Next = head
			head = head.Next
		}
	}
	return cur.Next
}

//func reverseListNode(head *ListNode, k int) {
//	first := ListNode{0, head}
//	head = head.Next
//	for i := 0; i < k-1; i++ {
//		head.Next = first.Next
//		first.Next = head
//		head = head.Next
//	}
//	return
//}

func getlength(head *ListNode) int {
	len := 0
	for head.Next != nil {
		len++
		head = head.Next
	}
	return len
}

func NewListNode(i int) *ListNode {
	return &ListNode{i, nil}
}

func main() {

	l := *NewListNode(0)
	for i := 1; i < 11; i++ {
		l = *NewListNode(i)
	}
	l = *reverseKGroup(&l, 2)
	for l.Next != nil {
		fmt.Println(l.Val)
	}
}

var RowSize = 100

type row struct {
	teacher string
	record  int
}

func getNewRow() row {
	return row{teacher: "", record: 0}
}

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) int {
	Table := []row{}
	var row = getNewRow()

	//推送模型伪代码
	var sum = 0
	var rowIndex = 0
	// Filter
	for rowIndex < RowSize {
		row = Table[rowIndex]
		rowIndex++
		if row.teacher == "Tom" {
			// Aggregation
			sum += row.record
		}
	}
	return sum

}
