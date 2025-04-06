package mutable

type IMutable interface {
	IsMutable() bool
	Mut() any
	UnMut() any
	IsConstant() bool
	AsConstant() any
	SetConstancy(value bool) any
}

type AbstractMutable[Field any] struct {
	parent     Field
	isMutable  bool
	isConstant bool
}

func ConstructAbstractMutable[Field any](parent *Field, mutable bool, constant bool) *AbstractMutable[Field] {
	if constant && mutable {
		mutable = false
	}

	return &AbstractMutable[Field]{
		parent:     *parent,
		isMutable:  mutable,
		isConstant: constant,
	}
}

func (sample *AbstractMutable[Field]) IsMutable() bool {
	return sample.isMutable
}

func (sample *AbstractMutable[Field]) Mut() Field {
	if !sample.isConstant {
		sample.isMutable = true
	}
	return sample.parent
}

func (sample *AbstractMutable[Field]) UnMut() Field {
	if !sample.isConstant {
		sample.isMutable = false
	}
	return sample.parent
}

func (sample *AbstractMutable[Field]) IsConstant() bool {
	return sample.isConstant
}

func (sample *AbstractMutable[Field]) SetConstancy(value bool) Field {
	if value {
		sample.isMutable = false
	}
	sample.isConstant = value
	return sample.parent
}

func (sample *AbstractMutable[Field]) AsConstant() Field {
	sample.isConstant = true
	sample.isMutable = false
	return sample.parent
}
