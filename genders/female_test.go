package genders

import "testing"

func TestFemaleName(t *testing.T) {
	if ok := Female.Name() == female; !ok {
		t.Fatalf("%s != %s", Female.Name(), female)
	}
}
