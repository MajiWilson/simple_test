package dp

/**
300. 最长递增子序列
中等
相关标签
相关企业
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


示例 1：

输入：nums = [10,9,2,5,3,7,101,108 18]
10
9
2
2 5
2 3
2 3 7
2 3 7 101
2 3 7 101 108
2 3 7 18 108
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：
*/

/*
*
注意这里栈中内容不是最终合法的子串， 但是数量是一致的
*/
func lengthOfLIS(nums []int) int {
	// 使用一个栈来维护遍历中的低增子串
	incArray := make([]int, 0)
	for _, num := range nums {
		//如果大于栈顶，直接添加到末尾即可
		if len(incArray) == 0 || incArray[len(incArray)-1] < num {
			incArray = append(incArray, num)
		}
		//否则，找到第一个小于num的位置，下一个位置使用这个更小的num代替，即不会改变当前的长度， 同时放宽，也保留了num的性质，将来也许可以用的上。
		i := len(incArray) - 1
		for ; i >= 0; i-- {
			if incArray[i] < num {
				incArray[i+1] = num
				break
			}
		}
		if i == -1 {
			incArray[0] = num
		}
	}

	return len(incArray)
}
