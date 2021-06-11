package animalcrossing

import "testing"

func TestPigName(t *testing.T) {
	if ok := Pig.Name() == pig; !ok {
		t.Fatal("Pig.Name() != pig")
	}
}
