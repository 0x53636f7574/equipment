package types

import (
	"github.com/0x53636f7574/equipment/interfaces"
	"github.com/0x53636f7574/equipment/traits"
)

type Map[Key comparable, Value any] map[Key]Value

func (theMap Map[Key, Value]) Size() int {
	return len(theMap)
}

func (theMap Map[Key, Value]) Drop(key Key) {
	delete(theMap, key)
}

func (theMap Map[Key, Value]) Keys() interfaces.Iterable[Key] {
	keys := Array[Key]{}

	for key, _ := range theMap {
		keys = keys.Append(key)
	}
	return keys
}

func (theMap Map[Key, Value]) Values() interfaces.Iterable[Value] {
	values := Array[Value]{}

	for _, value := range theMap {
		values = values.Append(value)
	}
	return values
}

func (theMap Map[Key, Value]) ForEach(callable func(index any, item Value)) {
	for key, value := range theMap {
		callable(key, value)
	}
}

func (theMap Map[Key, Value]) Map(mapper func(index any, item Value) any) interfaces.Iterable[any] {
	result := Array[any]{}

	for key, item := range theMap {
		result = result.Append(mapper(key, item))
	}
	return result
}

func (theMap Map[Key, Value]) Where(predicate func(index any, item Value) bool) interfaces.Iterable[Value] {
	result := make(Array[Value], 0, theMap.Size())

	for key, item := range theMap {
		if predicate(key, item) {
			result = result.Append(item)
		}
	}
	return result
}

func (theMap Map[Key, Value]) FirstWhere(predicate func(index any, item Value) bool) Value {
	for key, item := range theMap {
		if predicate(key, item) {
			return item
		}
	}
	return traits.Empty[Value]()
}
