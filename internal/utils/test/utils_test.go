package utils_test

import (
	"fmt"

	"github.com/ozoncp/ocp-note-api/internal/utils"
)

func AllTest() {
	SplitSliceTest()
}

func SplitSliceTest() {

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

	fmt.Println("*Split Slice Test*")

	for _, example := range dataset {
		result := utils.SplitSlice(example.Input, example.BatchSize)

		if EqualSlices(result, example.Output) {
			fmt.Printf("Test passed (Input: %v, output: %v, batchSize: %v)\n", example.Input, result, example.BatchSize)
		} else {
			fmt.Printf("Test failed (Input: %v, expected output: %v, output: %v, batchSize: %v)\n", example.Input, example.Output, result, example.BatchSize)
		}
	}
}

func EqualSlices(first [][]int, second [][]int) bool {

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
