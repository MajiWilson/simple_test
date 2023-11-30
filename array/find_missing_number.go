package array

/**
268. 丢失的数字
简单
793
相关企业
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。



示例 1：

输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 2：
*/

func missingNumber(nums []int) int {
	// 使用集合也是可以的
	n := len(nums)
	newNums := make([]int, n+1)
	// 各自其位
	for i := 0; i < n; i++ {
		newNums[nums[i]] = nums[i]
	}
	// 查看缺失的位置
	for i := 0; i < n+1; i++ {
		if newNums[i] != i {
			return i
		}
	}

	return 0

}
