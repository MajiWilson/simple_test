package String

/**
 * 560. 和为 K 的子数组	给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
 * 输入：nums = [1,1,1], k = 2
 * 输出：2
 *
 * Created by Master SkyWalker
 * May the force be with you !
 * 2020/4/9 , 12:41
 */

/**
  前缀和 + 哈希表
*/
func subarraySum(nums []int, k int) int {
	count := 0
	prefixCount := make(map[int]int)
	// 特殊情况，方便后续统一处理
	prefixCount[0] = 1
	sum := 0
	for _, num := range nums {
		sum += num
		// 找到连续区间
		if curCount, ok := prefixCount[sum-k]; ok {
			count += curCount
		}
		// 维护前缀和信息
		prefixCount[sum] = prefixCount[sum] + 1
	}
	return count

}
