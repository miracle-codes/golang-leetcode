package main

import (
	"fmt"
	"leetcode2022/func/makeList"
	"leetcode2022/func/showList"
)

//*
//设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。
//val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。
//假设链表中的所有节点都是 0-index 的。
//
//在链表类中实现这些功能：
//
//get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
//addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
//addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
//addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
//deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/design-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//*/

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := makeList.MakeListNode(arr)
	showList.ShowList(list)
	fmt.Println(get(list, 10))
	list = addAtHead(list, 12)
	showList.ShowList(list)
	fmt.Println(get(list, 10))
}

func get(list *makeList.ListNode, index int) int {
	temp := list
	for i := 1; i < index; i++ {
		if temp.Next == nil {
			return -1
		}
		temp = temp.Next
	}
	return temp.Val

}

func addAtHead(list *makeList.ListNode, value int) *makeList.ListNode {
	list2 := list
	temp := makeList.ListNode{Val: value, Next: list2}
	return &temp

}
