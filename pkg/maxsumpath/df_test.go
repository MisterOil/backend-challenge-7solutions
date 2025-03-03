package maxsumpath

import "testing"

func TestMaxPathSumDP(t *testing.T) {
	data := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}

	maxSum := MaxSumPathDP(data)
	expected := 237
	if maxSum != expected {
		t.Errorf("Expected maximum path sum %d, got %d", expected, maxSum)
	}
}
