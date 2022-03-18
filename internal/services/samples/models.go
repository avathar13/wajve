package samples

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

type Sample struct {
	Text   string     `json:"text"`
	Number *BigNumber `json:"number"`
	Found  bool       `json:"found"`
	Type   string     `json:"type"`
}

type BigNumber struct {
	Value string
}

var (
	errBigNumberFormat  = fmt.Errorf("wrong big number format")
	errNotSupportedType = fmt.Errorf("type is not supported")
)

func (b *BigNumber) UnmarshalJSON(data []byte) error {
	var dto interface{}
	if err := json.Unmarshal(data, &dto); err != nil {
		return fmt.Errorf("json.Unmarshal.error: %w", err)
	}

	switch value := dto.(type) {
	case float64:
		b.Value = strconv.FormatFloat(value, 'f', -1, 64)
	case string:
		trimStrings := strings.Split(value, "e+")

		if len(trimStrings) != 2 {
			return fmt.Errorf("%w, length after trim: %d", errBigNumberFormat, len(trimStrings))
		}

		number, err := strconv.ParseInt(trimStrings[0], 10, 64)
		if err != nil {
			return fmt.Errorf("strconv.ParseInt.error: %w, %s", err, trimStrings[0])
		}

		power, err := strconv.ParseInt(trimStrings[1], 10, 64)
		if err != nil {
			return fmt.Errorf("strconv.ParseInt.error: %w, %s", err, trimStrings[1])
		}

		b.Value = big.NewInt(number).Exp(big.NewInt(number), big.NewInt(power), nil).String()
	default:
		return fmt.Errorf("%w, type: %s", errNotSupportedType, reflect.TypeOf(value).String())
	}

	return nil
}

func (b *BigNumber) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(b.Value)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal.error: %w", err)
	}

	return bytes, nil
}
