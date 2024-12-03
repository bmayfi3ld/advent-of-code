package parser

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// convert matrix of numbers
// Input string block example
//
//	data := `
//
// 3   4
// 4   3
// 2   5
// 1   3
// 3   9
// 3   3
func BlockToSlices(data string) ([]int, []int, error) {
	// Prepare slices for each column
	var col1, col2 []int

	// Split the string into lines and iterate over each line
	lines := strings.Split(strings.TrimSpace(data), "\n")
	for _, line := range lines {
		// Split each line into fields (columns)
		fields := strings.Fields(line)
		if len(fields) == 2 {
			// Convert strings to integers and append to respective slices
			val1, err1 := strconv.Atoi(fields[0])
			val2, err2 := strconv.Atoi(fields[1])
			if err1 == nil && err2 == nil {
				col1 = append(col1, val1)
				col2 = append(col2, val2)
			} else {
				// missing err2
				return nil, nil, errors.WithMessage(err1, "Error converting string to integer:")
			}
		}
	}

	return col1, col2, nil
}
