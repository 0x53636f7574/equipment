package number

import (
	"math/big"
)

type BigNumber = *big.Int
type BigDecimal = *big.Float

type Bit = big.Word

type NaturalField interface {
	int | int8 | int16 | int32 | int64
}

type RealField interface {
	float32 | float64
}

func NewBigNumber() BigNumber {
	return new(big.Int)
}

func NewBigDecimal() BigDecimal {
	return new(big.Float)
}
