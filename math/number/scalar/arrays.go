package scalar

import "sort"

type NaturalScalarArray []*NaturalScalar

type DecimalScalarArray []*DecimalScalar

func (arr NaturalScalarArray) Len() int {
	return len(arr)
}

func (arr NaturalScalarArray) Less(i, j int) bool {
	return arr[i].LessThan(arr[j])
}

func (arr NaturalScalarArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr DecimalScalarArray) Len() int {
	return len(arr)
}

func (arr DecimalScalarArray) Less(i, j int) bool {
	return arr[i].LessThan(arr[j])
}

func (arr DecimalScalarArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr NaturalScalarArray) Sort() {
	sort.Sort(arr)
}

func (arr DecimalScalarArray) Sort() {
	sort.Sort(arr)
}

func (arr NaturalScalarArray) Limits() (int, int, *NaturalScalar, *NaturalScalar) {
	minIndex, maxIndex := 0, 0

	for index := range arr {
		if arr[minIndex].GreaterThan(arr[index]) {
			minIndex = index
		}
		if arr[maxIndex].LessThan(arr[index]) {
			maxIndex = index
		}
	}
	return minIndex, maxIndex, arr[minIndex], arr[maxIndex]
}

func (arr DecimalScalarArray) Limits() (int, int, *DecimalScalar, *DecimalScalar) {
	minIndex, maxIndex := 0, 0

	for index := range arr {
		if arr[minIndex].GreaterThan(arr[index]) {
			minIndex = index
		}
		if arr[maxIndex].LessThan(arr[index]) {
			maxIndex = index
		}
	}
	return minIndex, maxIndex, arr[minIndex], arr[maxIndex]
}
