package comparable

type IComparable[Field any] interface {
	Equals(Field) bool

	LessThan(Field) bool
	GreaterThan(Field) bool

	LessThanOrEqual(Field) bool
	GreaterThanOrEqual(Field) bool
}
