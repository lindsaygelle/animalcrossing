package genders

import "testing"

func TestMaleName(t *testing.T) {
	if ok := Male.Name() == male; !ok {
		t.Fatalf("%s != %s", Male.Name(), male)
	}
}
