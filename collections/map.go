package collections

type Map[Key comparable, Value any] map[Key]Value
type MapTransformer[Key, Value any] func(index Key, item Value) any
type MapFilter[Key, Value any] func(index any, item Value) bool

func NewMap[Key comparable, Value any](theMap map[Key]Value) Map[Key, Value] {
	return theMap
}

func (theMap *Map[Key, Value]) Unwrap() map[Key]Value {
	return *theMap
}

func (theMap *Map[Key, Value]) Size() int {
	return len(*theMap)
}

func (theMap *Map[Key, Value]) Drop(key Key) {
	delete(*theMap, key)
}

func (theMap *Map[Key, Value]) Keys() Array[Key] {
	keys := Array[Key]{}

	for key, _ := range *theMap {
		keys.Append(key)
	}
	return keys
}

func (theMap *Map[Key, Value]) Values() Array[Value] {
	values := Array[Value]{}

	for _, value := range *theMap {
		values.Append(value)
	}
	return values
}

func (theMap *Map[Key, Value]) Transform(transformer MapTransformer[Key, Value]) Array[any] {
	result := Array[any]{}

	for key, item := range *theMap {
		result.Append(transformer(key, item))
	}
	return result
}

func (theMap *Map[Key, Value]) Where(predicate MapFilter[Key, Value]) Array[Value] {
	result := make(Array[Value], 0, theMap.Size())

	for key, item := range *theMap {
		if predicate(key, item) {
			result.Append(item)
		}
	}
	return result
}

func (theMap *Map[Key, Value]) FirstWhere(predicate MapFilter[Key, Value]) Value {
	var empty Value
	for key, item := range *theMap {
		if predicate(key, item) {
			return item
		}
	}
	return empty
}
