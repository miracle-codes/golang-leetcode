package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	//list2 :=[]int{1,2,3,4,5}
	//a,b:=1,2

	//mergeInBetween(list1,a,b,list2)
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	temp1 := list1
	temp2 := list1
	for i := 0; i < a-1; i++ {
		temp1 = temp1.Next
	}
	for j := 0; j < b; j++ {
		temp2 = temp2.Next
	}
	temp1.Next = list2
	temp3 := list2
	for temp3.Next != nil {
		temp3 = temp3.Next
	}
	temp3.Next = temp2.Next
	return list1
}
