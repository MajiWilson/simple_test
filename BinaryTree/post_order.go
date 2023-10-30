package BinaryTree

/**
145. 二叉树的后序遍历
*/

// 递归算法
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 迭代算法（此外还有两种是基于先序遍历，然后结果反转）
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		curNode             *TreeNode
		recentLyVisitedNode *TreeNode
		stack               []*TreeNode
		res                 []int
	)

	curNode = root
	for len(stack) > 0 || curNode != nil {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}
		peekNode := stack[len(stack)-1]
		// 看有孩子节点是不是已经访问过了
		if peekNode.Right == nil && peekNode.Right != recentLyVisitedNode {
			curNode = peekNode.Right
		} else {
			res = append(res, peekNode.Val)
			stack = stack[:len(stack)-1]
			recentLyVisitedNode = peekNode
		}
	}

	return res
}
