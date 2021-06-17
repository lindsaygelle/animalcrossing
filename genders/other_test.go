package genders

import "testing"

func TestOtherName(t *testing.T) {
	if ok := Other.Name() == other; !ok {
		t.Fatalf("%s != %s", Other.Name(), other)
	}
}
