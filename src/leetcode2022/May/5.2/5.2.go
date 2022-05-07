package main

import "leetcode2022/func/makeList"

func main() {

}

func detectCycle(head *makeList.ListNode) *makeList.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
