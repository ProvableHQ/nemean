package account

import (
	"testing"
)

func TestParseAddress(t *testing.T) {
	addr := "aleo1ag4alvc4g7d4apzgvr5f4jt44l0aezev2dx8m0klgwypnh9u5uxs42rclr"
	res, err := ParseAddress(addr)
	if err != nil {
		t.Fatal(err)
	}

	t.Errorf("%v", res)
}
