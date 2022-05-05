package makeList

type ListNode struct {
	Val  int
	Next *ListNode
}

//func main() {
//	one := MakeListNode([]int{1, 2, 3, 5, 6, 7, 7, 2, 34, 2, 1})
//	for one != nil {
//		if one.Next == nil {
//			fmt.Println(one.Value)
//		} else {
//			fmt.Printf("%d ==> ", one.Value)
//		}
//		one = one.Next
//	}
//}

func MakeListNode(nums []int) *ListNode {

	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}

	return res
}
