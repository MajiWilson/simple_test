package String

import "sort"

/**
 * 49. 字母异位词分组
 */

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

// MAp
func groupAnagrams(strs []string) [][]string {
	resMap := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortStr := string(s)
		resMap[sortStr] = append(resMap[sortStr], str)
	}
	res := make([][]string, 0, len(resMap))
	for _, v := range resMap {
		res = append(res, v)
	}
	return res

}
