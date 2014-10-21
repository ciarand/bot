package test

import (
	"reflect"
	"testing"
)

// Equals errors if the two values (e - expected and a - actual) are different
func Equals(t *testing.T, e interface{}, a interface{}) {
	// if they're equal, we're done. no need to raise an error
	if reflect.TypeOf(e) != reflect.TypeOf(a) {
		t.Errorf("Expected %+v, got %+v", reflect.TypeOf(e), reflect.TypeOf(a))
	}

	if !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %+v, got %+v", e, a)
	}
}

// Ok errors if the provided error is not nil
func Ok(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

// Assert errors if the provided bool was not true
func Assert(t *testing.T, b bool, msg ...string) {
	if !b {
		if len(msg) == 0 {
			t.Error("expected false to be true")
		} else {
			t.Error(msg[0])
		}
	}
}
