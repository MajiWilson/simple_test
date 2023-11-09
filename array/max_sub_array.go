package array

import "math"

/**
53. 最大子数组和
中等
6.4K
相关企业
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。
*/
func maxSubArray(nums []int) int {
	sum := 0
	maxSum := math.MinInt

	for i := 0; i < len(nums); i++ {
		if sum <= 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		maxSum = int(math.Max(float64(sum), float64(maxSum)))
	}

	return maxSum

}
