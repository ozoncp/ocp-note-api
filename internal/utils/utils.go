package utils

func SplitSlice(slice []int, batchSize int) [][]int {

	if batchSize == 0 {
		return nil
	}

	sliceOfBatches := make([][]int, len(slice)/batchSize+1)

	for i := 0; i < len(slice); i++ {
		sliceOfBatches[i/batchSize] = append(sliceOfBatches[i/batchSize], slice[i])
	}

	return sliceOfBatches
}
