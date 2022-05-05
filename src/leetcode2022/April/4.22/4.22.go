package main

import "fmt"

func main() {
	arr := []int{1, 23, 1, 5, 6, 34, 1}
	list := makeList(arr)
	for list != nil {
		if list.Next == nil {
			fmt.Printf("%d", list.Value)
		} else {
			fmt.Printf("%d ==>", list.Value)
		}
		list = list.Next
	}

}

type ListNode struct {
	Value int
	Next  *ListNode
}

func makeList(arr []int) *ListNode {

	if len(arr) == 0 {
		return nil
	}
	list := &ListNode{
		Value: arr[0],
	}
	temp := list
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{Value: arr[i]}
		temp = temp.Next
	}

	return list
}
