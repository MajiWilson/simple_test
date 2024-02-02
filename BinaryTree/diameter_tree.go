package BinaryTree

/*
*
543. 二叉树的直径
简单
1.4K
相关企业
给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。
*/
func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftLength := dfs(root.Left)
		rightLength := dfs(root.Right)
		// 注意这里不是leftlength + rightlength + 1, 长度增加
		if leftLength+rightLength > maxLength {
			maxLength = leftLength + rightLength
		}
		// 长度增加（向上）
		if leftLength > rightLength {
			return leftLength + 1
		} else {
			return rightLength + 1
		}
	}

	dfs(root)
	return maxLength
}
