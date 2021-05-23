package utils

import "github.com/ozoncp/ocp-note-api/core/note"

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

func SplitNoteSlice(slice []note.Note, batchSize int) [][]note.Note {

	if batchSize <= 0 || slice == nil {
		return nil
	}

	var numberOfBatches int

	if len(slice)%batchSize == 0 {
		numberOfBatches = len(slice) / batchSize
	} else {
		numberOfBatches = len(slice)/batchSize + 1
	}

	sliceOfBatches := make([][]note.Note, 0, numberOfBatches)

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

func FilterSlice(slice []uint, filter []uint) []uint {
	var result []uint

	for _, val := range slice {
		if !containsValueInSlice(filter, val) {
			result = append(result, val)
		}
	}

	return result
}

func containsValueInSlice(slice []uint, value uint) bool {
	uniqValues := make(map[uint]struct{}, len(slice))

	for _, val := range slice {
		uniqValues[val] = struct{}{}
	}

	_, found := uniqValues[value]
	return found
}
