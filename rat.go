package sqlbigrat

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

var (
	_ sql.Scanner      = (*Rat)(nil)
	_ driver.Valuer    = (*Rat)(nil)
	_ json.Marshaler   = (*Rat)(nil)
	_ json.Unmarshaler = (*Rat)(nil)
)

type Rat struct {
	*big.Rat
	decimal.Decimal
	Valid bool
}

func (r *Rat) Scan(src any) error {
	if rat, ok := src.(*big.Rat); ok {
		r.Valid = true
		r.Rat = rat
		r.Decimal = decimal.NewFromBigRat(r.Rat, 32)
		return nil
	}

	switch val := src.(type) {
	case int:
		r.Decimal = decimal.NewFromInt(int64(val))
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case int8:
		r.Decimal = decimal.NewFromInt(int64(val))
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case int16:
		r.Decimal = decimal.NewFromInt(int64(val))
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case int32:
		r.Decimal = decimal.NewFromInt(int64(val))
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case int64:
		r.Decimal = decimal.NewFromInt(val)
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case uint:
		r.Decimal = decimal.NewFromInt(int64(val))
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case uint8:
		r.Decimal = decimal.NewFromInt(int64(val))
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case uint16:
		r.Decimal = decimal.NewFromInt(int64(val))
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case uint32:
		r.Decimal = decimal.NewFromInt(int64(val))
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case uint64:
		r.Decimal = decimal.NewFromInt(int64(val))
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case float32:
		r.Decimal = decimal.NewFromFloat32(val)
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case float64:
		r.Decimal = decimal.NewFromFloat(val)
		r.Rat = r.Decimal.Rat()
		r.Valid = true
		return nil
	case string:
		var ok bool
		r.Rat, ok = (&big.Rat{}).SetString(val)
		if !ok {
			return fmt.Errorf("[sqlrat.BqRat] invalid BIGNUMERIC value %+v of type %T", src, src)
		}
		r.Valid = true
		r.Decimal = decimal.NewFromBigRat(r.Rat, 32)
		return nil
	default:
		return errors.New(fmt.Sprintf("[sqlrat.BqRat] failed to convert value %+v to byte array of type %T", src, src))
	}
}

func (r *Rat) Value() (driver.Value, error) {
	if !r.Valid {
		return nil, nil
	}
	return r.Rat.String(), nil
}

func (r *Rat) MarshalJSON() ([]byte, error) {
	return r.Decimal.MarshalJSON()
}

func (r *Rat) UnmarshalJSON(data []byte) error {
	dec := decimal.NewFromInt(0)
	if err := dec.UnmarshalJSON(data); err != nil {
		return err
	}
	r.Decimal = dec
	r.Rat = dec.Rat()
	r.Valid = true
	return nil
}

func (r *Rat) String() string {
	return r.Decimal.String()
}
