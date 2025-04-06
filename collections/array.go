package collections

type Array[Covered any] []Covered
type ArrayFilter[Covered any] func(index int, item Covered) bool
type ArrayTransformer[Covered any] func(index int, item Covered) any

func NewArray[Covered any](array []Covered) Array[Covered] {
	return array
}

func (arr *Array[Covered]) Unwrap() []Covered {
	return *arr
}

func (arr *Array[Covered]) Append(item Covered) {
	*arr = append(*arr, item)
}

func (arr *Array[Covered]) DropAt(index int) {
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
}

func (arr *Array[Covered]) Empty() bool {
	return len(*arr) == 0
}

func (arr *Array[Covered]) Size() int {
	return len(*arr)
}

func (arr *Array[Covered]) Capacity() int {
	return cap(*arr)
}

func (arr *Array[Covered]) At(index int) Covered {
	var empty Covered
	if index >= arr.Size() {
		return empty
	}

	if index < 0 {
		index *= -1
		if index >= arr.Size() {
			return empty
		}
		index = arr.Size() - index
	}

	return (*arr)[index]
}

func (arr *Array[Covered]) First() Covered {
	return (*arr)[0]
}

func (arr *Array[Covered]) Last() Covered {
	return (*arr)[arr.Size()-1]
}

func (arr *Array[Covered]) Transform(transformer ArrayTransformer[Covered]) Array[any] {
	result := Array[any]{}

	for index, item := range *arr {
		result.Append(transformer(index, item))
	}
	return result
}

func (arr *Array[Covered]) Where(predicate ArrayFilter[Covered]) Array[Covered] {
	result := make(Array[Covered], 0, arr.Size())

	for index, item := range *arr {
		if predicate(index, item) {
			result.Append(item)
		}
	}
	return result
}

func (arr *Array[Covered]) FirstWhere(predicate ArrayFilter[Covered]) Covered {
	var empty Covered
	for index, item := range *arr {
		if predicate(index, item) {
			return item
		}
	}
	return empty
}
