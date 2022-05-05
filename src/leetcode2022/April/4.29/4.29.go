package main

import (
	"fmt"
	"leetcode2022/func/makeList"
	"leetcode2022/func/showList"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := makeList.MakeListNode(arr)
	showList.ShowList(list)
	list = swapPairs(list)
	fmt.Println("交换后")
	showList.ShowList(list)
}

func swapPairs(head *makeList.ListNode) *makeList.ListNode {
	dummyHead := &makeList.ListNode{0, head}
	temp := dummyHead

	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}

	return dummyHead.Next

}
