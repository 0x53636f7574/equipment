package interfaces

type Iterable[Covered any] interface {
	ForEach(func(Covered))

	Map(func(Covered) any) Iterable[any]

	Where(func(Covered) bool) Iterable[Covered]

	FirstWhere(func(Covered) bool) Covered
}
