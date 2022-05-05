package main

import (
	"leetcode2022/func/makeList"
	"leetcode2022/func/showList"
)

func main() {
	arr := []int{1, 23, 4, 5, 6, 7, 8, 9, 11}
	list := makeList.MakeListNode(arr)
	showList.ShowList(list)
	removeElements(list, 4)
	showList.ShowList(list)

}

func removeElements(head *makeList.ListNode, val int) *makeList.ListNode {
	if head == nil {
		return head
	}

	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head

}
