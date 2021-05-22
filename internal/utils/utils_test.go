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

		if equalSliceOfSlices(result, example.Output) {
			t.Logf("Test passed (Input: %v, output: %v, batchSize: %v)\n", example.Input, result, example.BatchSize)
		} else {
			t.Errorf("Test failed (Input: %v, expected output: %v, output: %v, batchSize: %v)\n", example.Input, example.Output, result, example.BatchSize)
		}
	}
}

func equalSliceOfSlices(first [][]uint, second [][]uint) bool {

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

func TestSwapKeyAndValue(t *testing.T) {
	type Dataset struct {
		Input  map[uint]string
		Output map[string]uint
	}

	dataset := []Dataset{
		{
			Input:  map[uint]string{1: "hex", 2: "mex", 3: "lex", 4: "zhex", 5: "chpex"},
			Output: map[string]uint{"hex": 1, "mex": 2, "lex": 3, "zhex": 4, "chpex": 5},
		},
		{
			Input:  nil,
			Output: nil,
		},
		{
			Input:  map[uint]string{1: "lex", 2: "lex", 3: "lex", 4: "zhex", 5: "chpex"},
			Output: nil,
		},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Logf("The test did not panic")
		} else {
			t.Logf("The test did panic")
		}
	}()

	for _, example := range dataset {
		result := SwapKeyAndValue(example.Input)

		if equalMaps(result, example.Output) {
			t.Logf("Test passed (Input: %v, output: %v)\n", example.Input, result)
		} else {
			t.Errorf("Test failed (Input: %v, expected output: %v, output: %v)\n", example.Input, example.Output, result)
		}
	}
}

func equalMaps(first map[string]uint, second map[string]uint) bool {

	if len(first) != len(second) {
		return false
	}

	for key, val := range first {
		if valFromSecond, found := second[key]; !found {
			return false
		} else {
			if val != valFromSecond {
				return false
			}
		}
	}

	return true
}

func TestFilterSlice(t *testing.T) {
	type Dataset struct {
		Input  []uint
		Filter []uint
		Output []uint
	}

	dataset := []Dataset{
		{
			Input:  []uint{1, 2, 3, 4, 5},
			Filter: []uint{1, 2},
			Output: []uint{3, 4, 5},
		},
		{
			Input:  []uint{1, 2, 3, 4, 5},
			Filter: []uint{11, 22},
			Output: []uint{1, 2, 3, 4, 5},
		},
		{
			Input:  nil,
			Filter: []uint{11, 22},
			Output: nil,
		},
		{
			Input:  []uint{1, 2, 3, 4, 5},
			Filter: nil,
			Output: []uint{1, 2, 3, 4, 5},
		},
		{
			Input:  nil,
			Filter: nil,
			Output: nil,
		},
		{
			Input:  []uint{1, 2, 3},
			Filter: []uint{1, 2, 3, 4, 5},
			Output: nil,
		},
	}

	for _, example := range dataset {
		result := FilterSlice(example.Input, example.Filter)

		if equalSlices(result, example.Output) {
			t.Logf("Test passed (Input: %v, output: %v, filter: %v)\n", example.Input, result, example.Filter)
		} else {
			t.Errorf("Test failed (Input: %v, expected output: %v, output: %v, filter: %v)\n", example.Input, example.Output, result, example.Filter)
		}
	}
}

func equalSlices(first []uint, second []uint) bool {

	if len(first) != len(second) {
		return false
	}

	for i, val := range first {
		if val != second[i] {
			return false
		}
	}

	return true
}
