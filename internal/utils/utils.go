package utils

func SplitSlice(slice []uint, batchSize int) [][]uint {

	if batchSize <= 0 || slice == nil {
		return nil
	}

	var numberOfBatches int

	if len(slice)%batchSize == 0 {
		numberOfBatches = len(slice) / batchSize
	} else {
		numberOfBatches = len(slice)/batchSize + 1
	}

	sliceOfBatches := make([][]uint, 0, numberOfBatches)

	for i := 0; i < len(slice); {
		end := i + batchSize

		if end > len(slice) {
			end = len(slice)
		}

		sliceOfBatches = append(sliceOfBatches, slice[i:end])
		i = end
	}

	return sliceOfBatches
}

func SwapKeyAndValue(data map[uint]string) map[string]uint {

	if data == nil {
		return nil
	}

	modifiedData := make(map[string]uint)

	for key, val := range data {
		if _, found := modifiedData[val]; found {
			panic("Key \"" + val + "\" already exists")
		}

		modifiedData[val] = key
	}

	return modifiedData
}
