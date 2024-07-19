package sqlbigrat

import (
	"database/sql/driver"
	"fmt"
	"math/big"
)

type Rat struct {
	*big.Rat
}

func NewFromInt(i int64) *Rat     { return &Rat{new(big.Rat).SetInt64(i)} }
func NewFromUInt(u uint64) *Rat   { return &Rat{new(big.Rat).SetUint64(u)} }
func NewFromFloat(f float64) *Rat { return &Rat{new(big.Rat).SetFloat64(f)} }

func NewFromString(s string) (*Rat, error) {
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		return nil, fmt.Errorf("[sqlbigrat.Rat] failed to convert value %q to byte array of type %T", s, s)
	}
	return &Rat{r}, nil
}

func (r *Rat) Scan(src any) error {
	if rat, ok := src.(*big.Rat); ok {
		r.Rat = rat
		return nil
	}

	switch val := src.(type) {
	case int:
		r.Rat = new(big.Rat).SetInt64(int64(val))
	case int8:
		r.Rat = new(big.Rat).SetInt64(int64(val))
	case int16:
		r.Rat = new(big.Rat).SetInt64(int64(val))
	case int32:
		r.Rat = new(big.Rat).SetInt64(int64(val))
	case int64:
		r.Rat = new(big.Rat).SetInt64(val)
	case uint:
		r.Rat = new(big.Rat).SetUint64(uint64(val))
	case uint8:
		r.Rat = new(big.Rat).SetUint64(uint64(val))
	case uint16:
		r.Rat = new(big.Rat).SetUint64(uint64(val))
	case uint32:
		r.Rat = new(big.Rat).SetUint64(uint64(val))
	case uint64:
		r.Rat = new(big.Rat).SetUint64(val)
	case float32:
		r.Rat = new(big.Rat).SetFloat64(float64(val))
	case float64:
		r.Rat = new(big.Rat).SetFloat64(val)
	case string:
		rat, ok := new(big.Rat).SetString(val)
		if !ok {
			return fmt.Errorf("[sqlbigrat.Rat] invalid value %+v of type %T", src, src)
		}
		r.Rat = rat
	default:
		return fmt.Errorf("[sqlbigrat.Rat] failed to convert value %+v to byte array of type %T", src, src)
	}

	return nil
}

func (r *Rat) Value() (driver.Value, error) {
	if r == nil || r.Rat == nil {
		return nil, nil
	}
	return r.Rat.String(), nil
}

func (r *Rat) MarshalJSON() ([]byte, error) {
	if r.Rat == nil {
		return []byte("null"), nil
	}
	return []byte(r.Rat.FloatString(16)), nil
}

func (r *Rat) UnmarshalJSON(data []byte) error {
	return r.Scan(string(data))
}

func (r *Rat) String() string {
	return r.Rat.FloatString(16)
}
