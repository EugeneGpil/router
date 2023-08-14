package methods

import (
	"bytes"
	"reflect"
	"testing"
)

func AssertSame(var1 interface{}, var2 interface{}, t *testing.T) {
	if reflect.TypeOf(var1) != reflect.TypeOf(var2) {
		t.Fatalf(`types do not match %T, %T`, var1, var2)
	}

	if reflect.TypeOf(var1).String() == "[]uint8" {
		AssertSame(bytes.Compare(var1.([]byte), var2.([]byte)), 0, t)

		return
	}

	if var1 != var2 {
		t.Fatalf(`var1 = %q, want match for %q`, var1, var2)
	}
}
