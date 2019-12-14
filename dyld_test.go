package dyld

import (
	"testing"
)

func TestAddrZero(t *testing.T) {
	_, err := Addr(0)
	if err != nil {
		return
	}
	t.Error("expected non-nil error")
}
