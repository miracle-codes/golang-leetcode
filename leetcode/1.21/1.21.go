package main

import "fmt"

func main() {
	Constructor()
}

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
	arr := []int{1, 1, 1, 12, 3, 4, 5, 6, 7}
	res := makeMyLinkedList(arr)
	for res != nil {
		fmt.Printf("%v ", res.val)
		res = res.next
	}
	return *res

}

func makeMyLinkedList(nums []int) *MyLinkedList {

	if len(nums) == 0 {
		return nil
	}

	res := &MyLinkedList{
		val: nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.next = &MyLinkedList{val: nums[i]}
		temp = temp.next
	}

	return res
}

func (this *MyLinkedList) Get(index int) int {
	return 0
}

func (this *MyLinkedList) AddAtHead(val int) {

}

func (this *MyLinkedList) AddAtTail(val int) {

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

}
