package animalcrossing

import "testing"

func TestBearName(t *testing.T) {
	if ok := Bear.Name() == bear; !ok {
		t.Fatal("Bear.Name() != bear")
	}
}
