package animalcrossing

import "testing"

func TestDodoName(t *testing.T) {
	if ok := Dodo.Name() == dodo; !ok {
		t.Fatal("Dodo.Name() != dodo")
	}
}
