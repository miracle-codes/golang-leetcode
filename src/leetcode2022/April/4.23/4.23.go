package main

import "fmt"
import "leetcode2022/func/makeList"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	list := makeList.MakeListNode(arr)
	for list != nil {
		if list.Next == nil {
			fmt.Println(list.Val)
		} else {
			fmt.Printf("%d ==>", list.Val)
		}
		list = list.Next
	}

}

//type ListNode struct {
//	Value int
//	Next *ListNode
//}

//func makeList(arr []int) *ListNode  {
//	if len(arr)==0 {
//		return nil
//	}
//	list:=&ListNode{
//		Value: arr[0],
//	}
//	temp := list
//	for i:=1;i<len(arr);i++{
//
//		temp.Next = &ListNode{
//			Value: arr[i],
//		}
//		temp = temp.Next
//	}
//
//	return list
//
//}
