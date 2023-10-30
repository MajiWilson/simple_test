package BinaryTree

/**
96. 不同的二叉搜索树
*/

// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

/**
动态规划思路：
	假设 n 个节点存在二叉排序树的个数是 G (n)，令 f(i) 为以 i 为根的二叉搜索树的个数，则
		G(n) =f(1)+f(2)+f(3)+f(4)+...+f(n)
		G(n) =G(0)∗G(n−1)+G(1)∗(n−2)+...+G(n−1)∗G(0)
  实际上这个数学上也叫做：卡特兰数

*/
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]

}
