package account

import (
	"testing"
)

func TestParseViewKey(t *testing.T) {
	k := "AViewKey1m8gvywHKHKfUzZiLiLoHedcdHEjKwo5TWo6efz8gK7wF"
	res, err := ParseViewKey(k)
	if err != nil {
		t.Fatal(err)
	}

	t.Errorf("%v", res)
}
