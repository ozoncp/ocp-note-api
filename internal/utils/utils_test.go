package utils

import (
	"testing"
)

func TestSplitSlice(t *testing.T) {

	type Dataset struct {
		BatchSize int
		Input     []uint
		Output    [][]uint
	}

	dataset := []Dataset{
		{
			BatchSize: 2,
			Input:     []uint{1, 2, 3, 4, 5, 6},
			Output:    [][]uint{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			BatchSize: 0,
			Input:     []uint{1, 2, 3, 4, 5, 6},
			Output:    nil,
		},
		{
			BatchSize: -6,
			Input:     []uint{1, 2, 3, 4, 5, 6},
			Output:    nil,
		},
		{
			BatchSize: 200,
			Input:     []uint{1, 2, 3, 4, 5, 6},
			Output:    [][]uint{{1, 2, 3, 4, 5, 6}},
		},
		{
			BatchSize: 6,
			Input:     []uint{1, 2, 3, 4, 5, 6},
			Output:    [][]uint{{1, 2, 3, 4, 5, 6}},
		},
		{
			BatchSize: 2,
			Input:     nil,
			Output:    nil,
		},
	}

	for _, example := range dataset {
		result := SplitSlice(example.Input, example.BatchSize)

		if equalMaps(result, example.Output) {
			t.Logf("Test passed (Input: %v, output: %v, batchSize: %v)\n", example.Input, result, example.BatchSize)
		} else {
			t.Errorf("Test failed (Input: %v, expected output: %v, output: %v, batchSize: %v)\n", example.Input, example.Output, result, example.BatchSize)
		}
	}
}

func equalMaps(first [][]uint, second [][]uint) bool {

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
