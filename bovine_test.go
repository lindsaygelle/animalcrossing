package animalcrossing

import "testing"

func TestBovineName(t *testing.T) {
	if ok := Bovine.Name() == bovine; !ok {
		t.Fatal("Bovine.Name() != bovine")
	}
}
