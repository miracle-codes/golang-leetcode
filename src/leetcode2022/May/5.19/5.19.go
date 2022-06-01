package main

type node struct {
	Val   int
	left  *node
	right *node
}

func main() {

}

func preorderTraversal(node2 *node) (res []int) {
	var traversal func(node *node)
	traversal = func(node *node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.left)
		traversal(node.right)

	}
	return res
}
