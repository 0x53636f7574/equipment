package scalar

import (
	"github.com/0x53636f7574/equipment/math/number"
	"github.com/0x53636f7574/equipment/traits/mutable"
	"strings"
)

type NaturalScalar struct {
	*mutable.AbstractMutable[*NaturalScalar]
	instance number.BigNumber
}

func NewNaturalScalar[Sample number.NaturalField | number.BigNumber | *NaturalScalar | *DecimalScalar | string](arg Sample) *NaturalScalar {
	scalar := &NaturalScalar{
		instance: number.NewBigNumber(),
	}
	scalar.AbstractMutable = mutable.ConstructAbstractMutable(&scalar, false, false)

	switch value := any(arg).(type) {
	case string:
		if strings.HasPrefix(value, "0x") {
			scalar.instance.SetString(strings.TrimLeft(value, "0x"), 16)
		} else {
			scalar.instance.SetString(value, 10)
		}
	case int:
		scalar.instance.SetInt64(int64(value))
	case int8:
		scalar.instance.SetInt64(int64(value))
	case int16:
		scalar.instance.SetInt64(int64(value))
	case int32:
		scalar.instance.SetInt64(int64(value))
	case int64:
		scalar.instance.SetInt64(value)
	case number.BigNumber:
		scalar.instance.Set(value)
	case *NaturalScalar:
		if value != nil && value.instance != nil {
			scalar.instance.Set(value.instance)
		}
	case *DecimalScalar:
		if value != nil && value.instance != nil {
			naturalValue, _ := value.instance.Int64()
			scalar.instance.SetInt64(naturalValue)
		}
	}
	return scalar
}

func (scalar *NaturalScalar) Bits() []number.Bit {
	return scalar.instance.Bits()
}

func (scalar *NaturalScalar) BitLen() int {
	return scalar.instance.BitLen()
}

func (scalar *NaturalScalar) Clone() *NaturalScalar {
	scalarCopy := &NaturalScalar{
		instance: number.NewBigNumber().Set(scalar.instance),
	}

	scalarCopy.AbstractMutable = mutable.ConstructAbstractMutable(&scalarCopy, scalar.IsMutable(), false)

	return scalarCopy
}

func (scalar *NaturalScalar) ToDecimalScalar() *DecimalScalar {
	decimalScalar := &DecimalScalar{
		instance: number.NewBigDecimal().SetInt(scalar.instance),
	}
	decimalScalar.AbstractMutable = mutable.ConstructAbstractMutable(&decimalScalar, scalar.IsMutable(), scalar.IsConstant())

	return decimalScalar
}

func (scalar *NaturalScalar) ToInt64() int64 {
	return scalar.instance.Int64()
}

func (scalar *NaturalScalar) ToUInt64() uint64 {
	return scalar.instance.Uint64()
}

func (scalar *NaturalScalar) moveOrCopy() *NaturalScalar {
	if scalar.IsConstant() || !scalar.IsMutable() {
		return scalar.Clone()
	}
	return scalar
}

func (scalar *NaturalScalar) Add(arg *NaturalScalar) *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.Add(result.instance, arg.instance)
	return result
}

func (scalar *NaturalScalar) Subtract(arg *NaturalScalar) *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.Sub(result.instance, arg.instance)
	return result
}

func (scalar *NaturalScalar) Multiply(arg *NaturalScalar) *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.Mul(result.instance, arg.instance)
	return result
}

func (scalar *NaturalScalar) Divide(arg *NaturalScalar) *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.Quo(result.instance, arg.instance)
	return result
}

func (scalar *NaturalScalar) Reminder(arg *NaturalScalar) *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.Rem(result.instance, arg.instance)
	return result
}

func (scalar *NaturalScalar) Pow(power int64) *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.Exp(result.instance, number.NewBigNumber().SetInt64(power), nil)
	return result
}

func (scalar *NaturalScalar) Sqrt() *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.Sqrt(result.instance)
	return result
}

func (scalar *NaturalScalar) Mod(arg *NaturalScalar) *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.Mod(result.instance, arg.instance)
	return result
}

func (scalar *NaturalScalar) Neg() *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.Neg(result.instance)
	return result
}

func (scalar *NaturalScalar) And(arg *NaturalScalar) *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.And(result.instance, arg.instance)
	return result
}

func (scalar *NaturalScalar) Or(arg *NaturalScalar) *NaturalScalar {
	result := scalar.moveOrCopy()
	result.instance.Or(result.instance, arg.instance)
	return result
}

func (scalar *NaturalScalar) Cmp(arg *NaturalScalar) int {
	return scalar.instance.Cmp(arg.instance)
}

func (scalar *NaturalScalar) Equals(arg *NaturalScalar) bool {
	return scalar.Cmp(arg) == 0
}

func (scalar *NaturalScalar) LessThan(arg *NaturalScalar) bool {
	return scalar.Cmp(arg) < 0
}

func (scalar *NaturalScalar) GreaterThan(arg *NaturalScalar) bool {
	return scalar.Cmp(arg) > 0
}

func (scalar *NaturalScalar) LessThanOrEqual(arg *NaturalScalar) bool {
	return scalar.Cmp(arg) <= 0
}

func (scalar *NaturalScalar) GreaterThanOrEqual(arg *NaturalScalar) bool {
	return scalar.Cmp(arg) >= 0
}

func (scalar *NaturalScalar) Increment() *NaturalScalar {
	return scalar.Add(NaturalOne)
}

func (scalar *NaturalScalar) Decrement() *NaturalScalar {
	return scalar.Subtract(NaturalOne)
}
