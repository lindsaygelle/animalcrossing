package animalcrossing

import "testing"

func TestAlligatorName(t *testing.T) {
	if ok := Alligator.Name() == alligator; !ok {
		t.Fatal("Alligator.Name() != alligator")
	}
}
