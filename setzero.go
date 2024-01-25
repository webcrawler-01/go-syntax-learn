//intital code for leetcode_03

func setZeroes(matrix [][]int) {
	//check zeroes in rowlevel
	//extra space

	var indexPositions [][]int = [][]int{} //empty go location
	//getting all the postions where the zero is located
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				indexPositions = append(indexPositions, []int{i, j})
			}
		}
	}

	//make the row and col to zero from that index
	for _, val := range indexPositions {
		for x := 0; x < len(matrix[0]); x++ {
			matrix[val[0]][x] = 0
		}
		for x := 0; x < len(matrix); x++ {
			matrix[x][val[1]] = 0
		}
	}
	// fmt.Println(indexPositions);
}