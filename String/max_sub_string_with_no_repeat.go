package String

import "math"

/**
3. 无重复字符的最长子串
中等
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

/**
双指针 + 哈希
*/

func lengthOfLongestSubstring(s string) int {
	left := 0
	chMap := make(map[int32]int)
	maxLength := 0
	for i, ch := range s {
		if idx, ok := chMap[ch]; ok {
			if left <= idx {
				left = idx + 1
			}
		}
		chMap[ch] = i
		maxLength = int(math.Max(float64(maxLength), float64(i-left+1)))
	}

	return maxLength

}
