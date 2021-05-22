package utils

func SplitSlice(slice []int, batchSize int) [][]int {

	if batchSize <= 0 || slice == nil {
		return nil
	}

	var numberOfBatches int

	if len(slice)%batchSize == 0 {
		numberOfBatches = len(slice) / batchSize
	} else {
		numberOfBatches = len(slice)/batchSize + 1
	}

	sliceOfBatches := make([][]int, numberOfBatches)

	for i := 0; i < len(slice); i++ {
		sliceOfBatches[i/batchSize] = append(sliceOfBatches[i/batchSize], slice[i])
	}

	return sliceOfBatches
}
