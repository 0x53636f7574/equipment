package scalar

import (
	"github.com/0x53636f7574/equipment/math/number"
	"github.com/0x53636f7574/equipment/traits/mutable"
)

type DecimalScalar struct {
	*mutable.AbstractMutable[*DecimalScalar]
	instance number.BigDecimal
}

func NewDecimalScalar[Sample number.RealField | *DecimalScalar | *NaturalScalar | string](arg Sample) *DecimalScalar {
	scalar := &DecimalScalar{
		instance: number.NewBigDecimal(),
	}
	scalar.AbstractMutable = mutable.ConstructAbstractMutable(&scalar, false, false)

	switch value := any(arg).(type) {
	case string:
		scalar.instance.SetString(value)
	case float32:
		scalar.instance.SetFloat64(float64(value))
	case float64:
		scalar.instance.SetFloat64(value)
	case *DecimalScalar:
		if value != nil && value.instance != nil {
			scalar.instance.Set(value.instance)
		}
	case *NaturalScalar:
		if value != nil && value.instance != nil {
			scalar.instance.SetInt(value.instance)
		}
	}

	return scalar
}

func (scalar *DecimalScalar) Clone() *DecimalScalar {
	scalarCopy := &DecimalScalar{
		instance: number.NewBigDecimal().Set(scalar.instance),
	}

	scalarCopy.AbstractMutable = mutable.ConstructAbstractMutable(&scalarCopy, scalar.IsMutable(), false)

	return scalarCopy
}

func (scalar *DecimalScalar) ToNaturalScalar(multiplier *NaturalScalar) *NaturalScalar {
	naturalValue := number.NewBigNumber()
	if multiplier == nil {
		multiplier = NewNaturalScalar(1)
	}
	NewDecimalScalar(multiplier).Mut().Multiply(scalar).UnMut().instance.Int(naturalValue)

	naturalScalar := &NaturalScalar{
		instance: naturalValue,
	}
	naturalScalar.AbstractMutable = mutable.ConstructAbstractMutable(&naturalScalar, scalar.IsMutable(), scalar.IsConstant())

	return naturalScalar
}

func (scalar *DecimalScalar) ToFloat64() float64 {
	result, _ := scalar.instance.Float64()
	return result
}

func (scalar *DecimalScalar) moveOrCopy() *DecimalScalar {
	if scalar.IsConstant() || !scalar.IsMutable() {
		return scalar.Clone()
	}
	return scalar
}

func (scalar *DecimalScalar) Add(arg *DecimalScalar) *DecimalScalar {
	result := scalar.moveOrCopy()
	result.instance.Add(result.instance, arg.instance)
	return result
}

func (scalar *DecimalScalar) Subtract(arg *DecimalScalar) *DecimalScalar {
	result := scalar.moveOrCopy()
	result.instance.Sub(result.instance, arg.instance)
	return result
}

func (scalar *DecimalScalar) Multiply(arg *DecimalScalar) *DecimalScalar {
	result := scalar.moveOrCopy()
	result.instance.Mul(result.instance, arg.instance)
	return result
}

func (scalar *DecimalScalar) Divide(arg *DecimalScalar) *DecimalScalar {
	result := scalar.moveOrCopy()
	result.instance.Quo(result.instance, arg.instance)
	return result
}

func (scalar *DecimalScalar) Sqrt() *DecimalScalar {
	result := scalar.moveOrCopy()
	result.instance.Sqrt(result.instance)
	return result
}

func (scalar *DecimalScalar) Neg() *DecimalScalar {
	result := scalar.moveOrCopy()
	result.instance.Neg(result.instance)
	return result
}

func (scalar *DecimalScalar) Cmp(arg *DecimalScalar) int {
	return scalar.instance.Cmp(arg.instance)
}

func (scalar *DecimalScalar) Equals(arg *DecimalScalar) bool {
	return scalar.Cmp(arg) == 0
}

func (scalar *DecimalScalar) LessThan(arg *DecimalScalar) bool {
	return scalar.Cmp(arg) < 0
}

func (scalar *DecimalScalar) GreaterThan(arg *DecimalScalar) bool {
	return scalar.Cmp(arg) > 0
}

func (scalar *DecimalScalar) LessThanOrEqual(arg *DecimalScalar) bool {
	return scalar.Cmp(arg) <= 0
}

func (scalar *DecimalScalar) GreaterThanOrEqual(arg *DecimalScalar) bool {
	return scalar.Cmp(arg) >= 0
}
