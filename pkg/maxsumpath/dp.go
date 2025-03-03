package maxsumpath

func MaxSumPathDP(data [][]int) int {
	if len(data) == 0 {
		return 0
	}
	depth := len(data)
	bottomArraySize := len(data[depth-1])

	bottomArray := make([]int, bottomArraySize) //<-- last row of data or base of the triangle
	copy(bottomArray, data[len(data)-1])

	// from bottom to top
	for i := depth - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			bottomArray[j] = data[i][j] + max(bottomArray[j], bottomArray[j+1])
		}
	}

	return bottomArray[0] // sum of the top element
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
