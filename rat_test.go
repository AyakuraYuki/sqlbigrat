package sqlbigrat

import (
	"encoding/json"
	"math/big"
	"testing"
)

func TestRat_MarshalJSON(t *testing.T) {
	r := Rat{Rat: big.NewRat(1, 3)}
	bs, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bs))
	t.Log(r.FloatString(16))
}
