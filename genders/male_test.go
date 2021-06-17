package genders

import "testing"

func TestMaleName(t *testing.T) {
	var s string = "Male"
	if ok := Male.Name() == s; !ok {
		t.Fatalf("%s != %s", Male.Name(), s)
	}
}
