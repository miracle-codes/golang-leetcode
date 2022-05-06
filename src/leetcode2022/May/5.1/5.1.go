package main

import (
	"leetcode2022/func/makeList"
	"leetcode2022/func/showList"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := makeList.MakeListNode(arr)
	list2 := list.Next.Next.Next

	list3 := &makeList.ListNode{1, list2}
	showList.ShowList(list)
	showList.ShowList(getIntersectionNode2(list, list3))
}

func getIntersectionNode(headA, headB *makeList.ListNode) *makeList.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func getIntersectionNode2(headA, headB *makeList.ListNode) *makeList.ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var step int
	var fast, slow *makeList.ListNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
