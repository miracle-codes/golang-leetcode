package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return head
	}

	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		if head.Next == nil {
			head.Val = 0
		}
		return head.Next
	}
	return head

}
