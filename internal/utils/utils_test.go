package utils_test

import ("fmt"
		"https://github.com/ozoncp/ocp-note-api/internal/utils")

func AllTest() {
	SplitSliceTest()
}

func SplitSliceTest() {

	type Dataset struct {
		BatchSize int
		Input []int
		Output [][]int
	}

	dataset := []Dataset {
		{ 2, []int{} { 1, 2, 3, 4, 5, 6 }, [][]int{} { {1, 2}, {3, 4}, {5, 6} } }
	}

	fmt.Println("*Split Slice Test*")

	
	for example := range dataset {
		if utils.SplitSlice(example.Input, example.BatchSize) != example.Output {
			fmt.Println("Test failed (Input: " + example.Input)
		} else {
			fmt.Println("Test passed (Input: " + example.Input)
		}
	}

}