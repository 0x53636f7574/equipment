package types

import (
	"github.com/0x53636f7574/equipment/interfaces"
	"github.com/0x53636f7574/equipment/traits"
)

type Array[Covered any] []Covered

func (arr Array[Covered]) Append(item Covered) Array[Covered] {
	return append(arr, item)
}

func (arr Array[Covered]) DropAt(index int) Array[Covered] {
	return append(arr[:index], arr[index+1:]...)
}

func (arr Array[Covered]) Empty() bool {
	return len(arr) == 0
}

func (arr Array[Covered]) Size() int {
	return len(arr)
}

func (arr Array[Covered]) Capacity() int {
	return cap(arr)
}

func (arr Array[Covered]) At(index int) Covered {
	if index >= arr.Size() {
		return traits.Empty[Covered]()
	}

	if index < 0 {
		index *= -1
		if index >= arr.Size() {
			return traits.Empty[Covered]()
		}
		index = arr.Size() - index
	}

	return arr[index]
}

func (arr Array[Covered]) First() Covered {
	return arr[0]
}

func (arr Array[Covered]) Last() Covered {
	return arr[arr.Size()-1]
}

func (arr Array[Covered]) ForEach(callable func(index any, item Covered)) {
	for index, item := range arr {
		callable(index, item)
	}
}

func (arr Array[Covered]) Map(mapper func(index any, item Covered) any) interfaces.Iterable[any] {
	result := Array[any]{}

	for index, item := range arr {
		result = result.Append(mapper(index, item))
	}
	return result
}

func (arr Array[Covered]) Where(predicate func(index any, item Covered) bool) interfaces.Iterable[Covered] {
	result := make(Array[Covered], 0, arr.Size())

	for index, item := range arr {
		if predicate(index, item) {
			result = result.Append(item)
		}
	}
	return result
}

func (arr Array[Covered]) FirstWhere(predicate func(index any, item Covered) bool) Covered {
	for index, item := range arr {
		if predicate(index, item) {
			return item
		}
	}
	return traits.Empty[Covered]()
}
