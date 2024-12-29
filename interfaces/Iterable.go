package interfaces

type Iterable[Covered any] interface {
	Size() int

	ForEach(func(any, Covered))

	Map(func(any, Covered) any) Iterable[any]

	Where(func(any, Covered) bool) Iterable[Covered]

	FirstWhere(func(any, Covered) bool) Covered
}
