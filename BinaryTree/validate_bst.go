package BinaryTree

import "math"

/**
98. 验证二叉搜索树


给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
*/

/**
错误：因为只保证了左右孩子合理不保证左右子树都合理！！！！！！！！！！！！！！
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
	}

	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right)

}

// 除了这个方法还有就是中序遍历判断有序（ 之所以使用入参而不是返回， 是因为bst是从上至下的）
func isValidBST2(root *TreeNode) bool {
	return judgeBst(root, math.MinInt, math.MaxInt)
}

func judgeBst(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return judgeBst(root.Left, lower, root.Val) && judgeBst(root.Right, root.Val, upper)

}
