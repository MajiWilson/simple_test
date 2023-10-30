package BinaryTree

/**
102. 二叉树的层序遍历
*/

// 借助队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		res [][]int
		q   []*TreeNode
	)
	q = append(q, root)

	for len(q) > 0 {
		levelSize := len(q)
		levelRes := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			levelRes[i] = q[i].Val
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		q = q[levelSize:]
		res = append(res, levelRes)
	}

	return res
}
