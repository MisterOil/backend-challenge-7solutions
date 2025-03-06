package leftrightequal

import (
	"fmt"
	"strings"
)

func MinDigitSequence(encoded string) (string, error) {

	encodedStrLength := len(encoded)
	var groups []Group
	var error error
	groupStart := 0
	for i := 0; i < encodedStrLength; i++ {
		if encoded[i] != '=' && encoded[i] != 'R' && encoded[i] != 'L' {
			return "", fmt.Errorf("invalid input")
		}
		if encoded[i] != '=' {
			groups = append(groups, Group{start: groupStart, end: i})
			groupStart = i + 1
		}
	}
	groups = append(groups, Group{start: groupStart, end: encodedStrLength})
	// fmt.Println(groups)

	groupsLength := len(groups)
	groupValues := make([]int, groupsLength)

	for i := 0; i < groupsLength-1; i++ {
		boundaryIndex := groups[i].end
		if encoded[boundaryIndex] == 'R' {
			groupValues[i+1] = groupValues[i] + 1
		}
	}

	for i := groupsLength - 2; i >= 0; i-- {
		boundaryIndex := groups[i].end
		if encoded[boundaryIndex] == 'L' {
			if groupValues[i] <= groupValues[i+1] {
				groupValues[i] = groupValues[i+1] + 1
			}
		}
	}

	digits := make([]string, encodedStrLength+1)
	for i, group := range groups {
		digitStr := fmt.Sprintf("%d", groupValues[i])
		for j := group.start; j <= group.end; j++ {
			digits[j] = digitStr
		}
	}

	result := strings.Join(digits, "")

	return result, error
}
