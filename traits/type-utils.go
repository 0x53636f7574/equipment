package traits

func Empty[T any]() T {
	var empty T
	return empty
}

func UnwrapFirst(items ...any) any {
	return items[0]
}

func OmitErrors(items ...any) []any {
	var result []any
	for _, item := range items {
		if _, impl := item.(error); !impl {
			result = append(result, item)
		}
	}
	return result
}
