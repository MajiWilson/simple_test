package BinaryTree

/**
94. 二叉树的中序遍历
*/

// 递归写法
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// 非递归写法： 借助stack
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		stack []*TreeNode
		node  *TreeNode
	)
	node = root
	stack = append(stack, node)
	node = node.Left
	res := make([]int, 0)

	// 有两个条件，防止提前结束遍历： 每个节点仅入栈一次，出栈一次
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node) // push
			root = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop
		res = append(res, node.Val)
		node = node.Right

	}
	return res
}

// 非递归写法： 借助stack 同2，优化了下顺序
func inorderTraversal3(root *TreeNode) []int {
	ret := make([]int, 0)
	var stack []*TreeNode
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		node = node.Right
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}

	return ret

}
