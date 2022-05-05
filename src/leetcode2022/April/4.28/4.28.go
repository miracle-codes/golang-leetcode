package main

import (
	"leetcode2022/func/makeList"
	"leetcode2022/func/showList"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := makeList.MakeListNode(arr)
	showList.ShowList(list)
	list = reverseList(list)
	showList.ShowList(list)
}

func reverseList(head *makeList.ListNode) *makeList.ListNode {
	var prev *makeList.ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev

}
