package String

/**
239. 滑动窗口最大值
提示
困难
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。
*/

// 维护一个单调队列即可（存储序列）， 注意最大值不一定还在窗口中
func maxSlidingWindow(nums []int, k int) []int {

	n := len(nums)
	dq := make([]int, 0)
	if n == 0 || n < k {
		return []int{}
	}
	res := make([]int, n-k+1)

	// 处理第一个窗口
	for i := 0; i < k; i++ {
		for len(dq) != 0 && nums[dq[len(dq)-1]] <= nums[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}
	res[0] = nums[dq[0]]

	// 处理其他窗口
	for i := k; i < n; i++ {
		for len(dq) != 0 && nums[dq[len(dq)-1]] <= nums[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		// 去调已经在窗口外的数据
		for dq[0] < i-k+1 {
			dq = dq[1:]
		}
		res[i-k+1] = nums[dq[0]]
	}

	return res
}
