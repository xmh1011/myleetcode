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
	fmt.Println(getFactorial(23+12-2) / (getFactorial(23-1) * getFactorial(12-1)))
	fmt.Println(getFactorial(23 + 12 - 2))
	fmt.Println(getFactorial(22))
	fmt.Println(getFactorial(12 - 1))

}

func getFactorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * getFactorial(n-1)
	}
}

var RowSize = 100
