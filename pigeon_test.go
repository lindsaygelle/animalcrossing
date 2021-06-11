package animalcrossing

import "testing"

func TestPigeonName(t *testing.T) {
	if ok := Pigeon.Name() == pigeon; !ok {
		t.Fatal("Pigeon.Name() != pigeon")
	}
}
