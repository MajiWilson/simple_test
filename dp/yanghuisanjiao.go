package dp

func generate(numRows int) [][]int {
	var res [][]int
	if numRows < 1 {
		return res
	}
	res = append(res, []int{1})

	for i := 1; i < numRows; i++ {
		newRow := make([]int, len(res[i-1])+1)
		for j := 0; j < len(res[i-1])+1; j++ {
			if j == 0 || j == len(res[i-1]) {
				newRow[j] = 1
			} else {
				newRow[j] = res[i-1][j] + res[i-1][j-1]
			}
		}
		res = append(res, newRow)
	}

	return res
}
