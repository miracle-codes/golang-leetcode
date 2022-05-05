package main

import (
	"leetcode2022/func/makeList"
	"leetcode2022/func/showList"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	list := makeList.MakeListNode(arr)
	showList.ShowList(list)
	list = removeNthFromEnd(list, 5)
	showList.ShowList(list)
}

func getLength(head *makeList.ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}

	return length
}

func removeNthFromEnd(head *makeList.ListNode, n int) *makeList.ListNode {
	length := getLength(head)
	dummy := &makeList.ListNode{0, head}
	cur := dummy
	if length < n {
		return nil
	}

	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next

	return dummy.Next
}

func removeNthFromEnd2(head *makeList.ListNode, n int) *makeList.ListNode {
	length := getLength(head)
	dummy := &makeList.ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next

}
