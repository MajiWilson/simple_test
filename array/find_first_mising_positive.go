package array

/**
41. 缺失的第一个正数
提示
困难
2K
相关企业
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。


示例 1：

输入：nums = [1,2,0]
输出：3
示例 2：

输入：nums = [3,4,-1,1]
输出：2

*/

/**
一个重要结论： 第一个缺失的值一定在1 - N+1之间
*/
func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 将负数设置成》= n+1的值
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}

		if num <= n {
			if nums[num-1] > 0 {
				nums[num-1] = -nums[num-1]
			}
		}
	}

	// 找第一个非负数位置索引即可
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1

}

/**
优化： 更加直观
*/

func firstMissingPositive2(nums []int) int {

	n := len(nums)
	// 原位哈希
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			temp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = temp
		}
	}

	// 各置其位， 不满足的即为第一个
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1

}
