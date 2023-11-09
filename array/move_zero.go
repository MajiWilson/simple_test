package array

/**
283. 移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

 [0] [1] 0 3 12
 1 [0] [0] 3 12
 1 [0] 0 [3] 12
 1 3 [0] 0 [12]
 1 3 12 [0] 0


[2] [3] 4 0 6
2 [3][4] 0 6
2 3 [4] [0] 6
2 3 4 [0] [6]
2 3 4 6 [0]
*/
func moveZeroes(nums []int) {
	left := 0
	right := 1
	// 0～left-1 之间的数据都是已经整理后的，right是遍历指针， left- right-1之间包含连续的0
	for right < len(nums) {
		if nums[left] == 0 && nums[right] != 0 {
			nums[left] = nums[right]
			nums[right] = 0
			left++
			right++
		} else if nums[left] == 0 && nums[right] == 0 {
			right++
		} else {
			left++
			right++
		}
	}

}
