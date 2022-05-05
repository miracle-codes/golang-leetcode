package main

import (
	"leetcode2022/func/makeList"
	"leetcode2022/func/showList"
)

func main() {
	arr := []int{1, 2, 6, 3, 4, 5, 6}
	list := makeList.MakeListNode(arr)
	showList.ShowList(list)
	removeElements(list, 6)
	showList.ShowList(list)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *makeList.ListNode, val int) *makeList.ListNode {
	if head == nil {
		return head
	}
	for head.Val == val {
		head.Val = 0
		head = head.Next
	}

	for head.Next != nil {
		if head.Next.Val == val {
			if head.Next.Next == nil {
				head.Next.Val = 0
			} else {
				head = head.Next.Next
			}
		}
		head = head.Next
	}
	return head

}
