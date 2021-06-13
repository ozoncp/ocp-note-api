package utils

import (
	"errors"

	"github.com/ozoncp/ocp-note-api/core/note"
)

func SplitSlice(data []uint, batchSize int) [][]uint {

	if batchSize <= 0 || data == nil {
		return nil
	}

	var numberOfBatches int

	if len(data)%batchSize == 0 {
		numberOfBatches = len(data) / batchSize
	} else {
		numberOfBatches = len(data)/batchSize + 1
	}

	sliceOfBatches := make([][]uint, 0, numberOfBatches)

	for i := 0; i < len(data); {
		end := i + batchSize

		if end > len(data) {
			end = len(data)
		}

		sliceOfBatches = append(sliceOfBatches, data[i:end])
		i = end
	}

	return sliceOfBatches
}

func SplitNoteSlice(data []note.Note, batchSize uint32) [][]note.Note {

	if batchSize <= 0 || data == nil {
		return nil
	}

	var numberOfBatches int

	if len(data)%int(batchSize) == 0 {
		numberOfBatches = len(data) / int(batchSize)
	} else {
		numberOfBatches = len(data)/int(batchSize) + 1
	}

	sliceOfBatches := make([][]note.Note, 0, numberOfBatches)

	for i := 0; i < len(data); {
		end := i + batchSize

		if end > len(data) {
			end = len(data)
		}

		sliceOfBatches = append(sliceOfBatches, data[i:end])
		i = end
	}

	return sliceOfBatches
}

func SwapKeyAndValue(data map[uint]string) map[string]uint {

	if data == nil {
		return nil
	}

	modifiedData := make(map[string]uint, len(data))

	for key, val := range data {
		if _, found := modifiedData[val]; found {
			panic("key \"" + val + "\" already exists")
		}

		modifiedData[val] = key
	}

	return modifiedData
}

func FilterSlice(data []uint, filter []uint) []uint {
	var result []uint

	uniqValues := make(map[uint]struct{}, len(data))

	for _, val := range filter {
		uniqValues[val] = struct{}{}
	}

	for _, val := range data {
		if !containsValueInSlice(uniqValues, val) {
			result = append(result, val)
		}
	}

	return result
}

func containsValueInSlice(data map[uint]struct{}, value uint) bool {
	_, found := data[value]
	return found
}

func ConvertSliceToMap(data []note.Note) (map[uint64]note.Note, error) {

	if len(data) == 0 {
		return nil, errors.New("the slice is empty")
	}

	modifiedData := make(map[uint64]note.Note, len(data))

	for _, val := range data {
		if _, found := modifiedData[val.Id]; found {
			return nil, errors.New("the key already exists")
		}

		modifiedData[val.Id] = val
	}

	return modifiedData, nil
}
