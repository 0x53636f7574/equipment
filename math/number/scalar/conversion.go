package scalar

import (
	"encoding/json"
	"errors"
	"github.com/0x53636f7574/equipment/math/number"
	"strings"
)

func (scalar *NaturalScalar) String() string {
	return scalar.instance.String()
}

func (scalar *NaturalScalar) MarshalJSON() ([]byte, error) {
	if scalar.instance == nil {
		return nil, nil
	}

	return []byte(scalar.String()), nil
}

func (scalar *NaturalScalar) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)

	if err != nil {
		return err
	}

	base := 10
	if strings.HasPrefix(value, "0x") {
		base = 16
	}

	if scalar.instance == nil {
		scalar.instance = number.NewBigNumber()
	}

	_, done := scalar.instance.SetString(value, base)

	if !done {
		return errors.New("string doesn't represent an integer number")
	}

	return nil
}

func (scalar *DecimalScalar) String() string {
	return scalar.instance.String()
}

func (scalar *DecimalScalar) MarshalJSON() ([]byte, error) {
	if scalar.instance == nil {
		return nil, nil
	}

	return []byte(scalar.String()), nil
}

func (scalar *DecimalScalar) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)

	if err != nil {
		return err
	}

	if scalar.instance == nil {
		scalar.instance = number.NewBigDecimal()
	}

	_, done := scalar.instance.SetString(value)

	if !done {
		return errors.New("string doesn't represent an number")
	}

	return nil
}
