package BinaryTree

/**
113. 路径总和 II
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。
*/

// DFS  + 回溯： 指针
func pathSum(root *TreeNode, targetSum int) [][]int {
	paths := make([][]int, 0)
	path := make([]int, 0)
	helper(root, targetSum, &paths, &path)
	return paths
}

func helper(root *TreeNode, targetSum int, paths *[][]int, curPath *[]int) {
	if root == nil {
		return
	}
	*curPath = append(*curPath, root.Val)

	// 叶子节点
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {

			var newPath []int
			for _, val := range *curPath {
				newPath = append(newPath, val)
			}
			*paths = append(*paths, newPath)
		}
	}

	helper(root.Left, targetSum-root.Val, paths, curPath)
	helper(root.Right, targetSum-root.Val, paths, curPath)
	*curPath = (*curPath)[:len(*curPath)-1]
	return
}

// DFS  + 回溯： 全局变量
func pathSum2(root *TreeNode, targetSum int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, target int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && target == node.Val {
			res = append(res, append([]int(nil), path...))
			return
		}
		dfs(node.Left, target-node.Val)
		dfs(node.Right, target-node.Val)
	}
	dfs(root, targetSum)
	return res
}
