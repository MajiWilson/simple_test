package array

import "math"

/**
 * 42
 * desc: 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 */

/**
  动态规划 ：
*/
func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	rightMax[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = int(math.Max(float64(leftMax[i-1]), float64(height[i])))
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = int(math.Max(float64(rightMax[i+1]), float64(height[i])))
	}

	// 计算
	maxVol := 0
	for i := 0; i < len(height); i++ {
		lowBorder := int(math.Min(float64(leftMax[i]), float64(rightMax[i])))
		maxVol += lowBorder - height[i]
	}

	return maxVol
}

/*
  单调栈
*/

func trap2(height []int) int {
	i := 0
	vol := 0
	stack := make([]int, 0)
	for i < len(height) {
		if len(stack) == 0 || height[stack[len(stack)-1]] > height[i] {
			stack = append(stack, i)
			i++
		} else {
			idxTop := stack[len(stack)-1]
			// pop
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				continue
			}
			vol += (int(math.Min(float64(height[i]), float64(height[stack[len(stack)-1]]))) - height[idxTop]) * (i - stack[len(stack)-1] - 1)
		}
	}
	return vol

}

/*
	双指针
*/
func trap3(height []int) int {
	return 0
}
