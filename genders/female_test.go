package genders

import "testing"

func TestFemaleName(t *testing.T) {
	var s string = "Female"
	if ok := Female.Name() == s; !ok {
		t.Fatalf("%s != %s", Female.Name(), s)
	}
}
