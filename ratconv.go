package sqlbigrat

import (
	"github.com/shopspring/decimal"
	"strconv"
)

// ToDecimal converts sqlbigrat.Rat into shopspring/decimal.Decimal, with given precision
func (r *Rat) ToDecimal(precision int32) decimal.Decimal {
	if r == nil {
		return decimal.NewFromInt(0)
	}
	return decimal.NewFromBigRat(r.Rat, precision)
}

// ToInt returns integer part of sqlbigrat.Rat
func (r *Rat) ToInt(base, bitSize int) (int64, error) {
	if r == nil {
		return 0, nil
	}
	num, err := strconv.ParseInt(r.Rat.FloatString(2), base, bitSize)
	return num, err
}
