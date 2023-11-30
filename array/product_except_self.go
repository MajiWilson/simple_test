package array

/**
238. 除自身以外数组的乘积
中等
1.6K
相关企业
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。



示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:
*/

func productExceptSelf(nums []int) []int {

	leftMultiply := make([]int, len(nums))
	rightMultiply := make([]int, len(nums))
	leftMultiply[0] = 1
	rightMultiply[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		leftMultiply[i] = leftMultiply[i-1] * nums[i-1]
	}
	for j := len(nums) - 2; j >= 0; j-- {
		rightMultiply[j] = rightMultiply[j+1] * nums[j+1]
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = leftMultiply[i] * rightMultiply[i]
	}
	return res
}
