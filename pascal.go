func generate(numRows int) [][]int {
	var solution [][]int = [][]int{}
	for i := 0; i < numRows; i++ {
		var innerRow []int = make([]int, i+1) // Increase size by 1
		for k := 0; k <= i; k++ {
			if k == 0 || k == len(innerRow)-1 {
				innerRow[k] = 1
			} else {
				innerRow[k] = solution[i-1][k] + solution[i-1][k-1]
			}
		}
		solution = append(solution, innerRow)
	}

	return solution
}
