package BinaryTree

/**
144. 二叉树的前序遍历
*/

// 递归算法
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

// 迭代算法：借助栈， 因为Root访问过就可以出栈所以比较简单
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		res   []int
		stack []*TreeNode
	)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		// 注意右孩子先入栈后访问
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}

// 递归算法： 借助栈， 一直遍历左孩子， 入栈前访问，
func preorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		res     []int
		stack   []*TreeNode
		curNode *TreeNode
	)
	curNode = root
	stack = append(stack, curNode)
	res = append(res, curNode.Val)
	curNode = curNode.Left

	for len(stack) > 0 || curNode != nil {
		for curNode != nil {
			stack = append(stack, curNode)
			res = append(res, curNode.Val)
			curNode = curNode.Left
		}
		peerNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curNode = peerNode.Right
	}
	return res
}
