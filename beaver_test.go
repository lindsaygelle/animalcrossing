package animalcrossing

import "testing"

func TestBeaverName(t *testing.T) {
	if ok := Beaver.Name() == beaver; !ok {
		t.Fatal("Beaver.Name() != beaver")
	}
}
