package BinaryTree

/**
107. 二叉树的层序遍历 II
*/
//给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

func levelOrderBottom(root *TreeNode) [][]int {
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

	// 逆转（注意序号边界）
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
