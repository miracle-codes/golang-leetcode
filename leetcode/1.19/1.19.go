package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {
	arr := []int{1, 2, 3, 4, 6, 5, 6}
	head := &ListNode{Value: arr[0]}
	tail := head //需要头尾两个指针
	for i := 1; i < len(arr); i++ {
		//方法一 数组直接构建链表
		tail.Next = &ListNode{Value: arr[i]} //tail.Next => (*tail).Next  golang的隐式转换
		tail = tail.Next
		//head.Append(list[i]) //方法二 追加构建链表
	}

	//one := makeListNode(arr)
	removeElements(head, 1)
	//for one != nil {
	//	fmt.Println(one.Value)
	//	one = one.Next
	//}
	head.Show()

}

func (h *ListNode) Append(i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode{Value: i}
}
func (h *ListNode) Show() {
	fmt.Println(h.Value)
	for h.Next != nil {
		h = h.Next
		fmt.Println(h.Value)
	}
}

func makeListNode(nums []int) *ListNode {

	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Value: nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Value: nums[i]}
		temp = temp.Next
	}

	return res
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Value == val {
		if head.Next == nil {
			head.Value = 0
		}
		return head.Next
	}
	return head

}
