package array

import "sort"

/**
15. 三数之和
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	//先排序
	sort.Ints(nums)

	// 先确定第一个数
	for i := 0; i < len(nums); i++ {
		// 无需判断
		if nums[i] > 0 {
			return res
		}
		// 避免重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 确定剩下的两个数（ 两数之和问题）
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				// 避免重复
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return res

}
