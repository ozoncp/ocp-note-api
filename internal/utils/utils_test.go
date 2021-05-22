package utils

import (
	"testing"
)

func TestSplitSlice(t *testing.T) {

	type Dataset struct {
		BatchSize int
		Input     []int
		Output    [][]int
	}

	dataset := []Dataset{
		{
			BatchSize: 2,
			Input:     []int{1, 2, 3, 4, 5, 6},
			Output:    [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			BatchSize: 0,
			Input:     []int{1, 2, 3, 4, 5, 6},
			Output:    nil,
		},
		{
			BatchSize: -6,
			Input:     []int{1, 2, 3, 4, 5, 6},
			Output:    nil,
		},
		{
			BatchSize: 200,
			Input:     []int{1, 2, 3, 4, 5, 6},
			Output:    [][]int{{1, 2, 3, 4, 5, 6}},
		},
		{
			BatchSize: 6,
			Input:     []int{1, 2, 3, 4, 5, 6},
			Output:    [][]int{{1, 2, 3, 4, 5, 6}},
		},
		{
			BatchSize: 2,
			Input:     nil,
			Output:    nil,
		},
	}

	for _, example := range dataset {
		result := SplitSlice(example.Input, example.BatchSize)

		if equalSlices(result, example.Output) {
			t.Logf("Test passed (Input: %v, output: %v, batchSize: %v)\n", example.Input, result, example.BatchSize)
		} else {
			t.Errorf("Test failed (Input: %v, expected output: %v, output: %v, batchSize: %v)\n", example.Input, example.Output, result, example.BatchSize)
		}
	}
}

func equalSlices(first [][]int, second [][]int) bool {

	if len(first) != len(second) {
		return false
	}

	for i, slice := range first {
		if len(slice) != len(second[i]) {
			return false
		}

		for j, val := range slice {
			if val != second[i][j] {
				return false
			}
		}
	}

	return true
}

func SwapKeyAndValueTest() {

	// type Dataset struct {
	// 	Input  [uint]string
	// 	Output [string]uint
	// }

	// dataset := []Dataset{[uint]string{1: "hex", 2: "mex"}, [string]uint{"hex"}}

	// dataset := []Dataset{
	// 	{
	// 		Input:  []int{1, 2, 3, 4, 5, 6},
	// 		Output: [][]int{{1, 2}, {3, 4}, {5, 6}},
	// 	},
	// 	{
	// 		Input:  []int{1, 2, 3, 4, 5, 6},
	// 		Output: [][]int{{1, 2}, {3, 4}, {5, 6}},
	// 	},
	// 	{
	// 		Input:  nil,
	// 		Output: nil,
	// 	},
	// }

}
