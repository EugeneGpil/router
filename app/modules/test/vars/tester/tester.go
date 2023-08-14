package tester

import (
	"bytes"
	"reflect"
	"testing"
)

var t *testing.T

func SetTester(tester *testing.T) {
	t = tester
}

func AssertSame(var1 interface{}, var2 interface{}) {
	if reflect.TypeOf(var1) != reflect.TypeOf(var2) {
		t.Fatalf(`types do not match %T, %T`, var1, var2)
	}

	if reflect.TypeOf(var1).String() == "[]uint8" {
		AssertSame(bytes.Compare(var1.([]byte), var2.([]byte)), 0)

		return
	}

	if var1 != var2 {
		t.Fatalf(`var1 = %q, want match for %q`, var1, var2)
	}
}

func AssertNotNil(var1 interface{}) {
	if var1 == nil {
		t.Fatalf(`variable is nil`)
	}
}

func AssertLen(var1 interface{}, length int) {
	if len := reflect.ValueOf(var1).Len(); len != length {
		t.Fatalf(`routes count = %q, want match for %q`, len, length)
	}
}
