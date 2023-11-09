package array

import "math"

/**
11. 盛最多水的容器
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。
*/
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := -1
	for left < right {
		if height[left] > height[right] {
			maxArea = int(math.Max(float64(height[right]*(right-left)), float64(maxArea)))
			right--
		} else {
			maxArea = int(math.Max(float64(height[left]*(right-left)), float64(maxArea)))
			left++
		}
	}
	return maxArea
}
