package cloneable

type ICloneable[Field any] interface {
	Clone() Field
}
