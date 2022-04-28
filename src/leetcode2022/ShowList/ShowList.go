package showList

import (
	"fmt"
	"leetcode2022/makeList"
)

func ShowList(node *makeList.ListNode) {
	for node != nil {
		if node.Next == nil {
			fmt.Println(node.Val)
		} else {
			fmt.Printf("%d ==>", node.Val)
		}
		node = node.Next

	}

}
