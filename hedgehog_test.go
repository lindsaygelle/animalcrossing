package animalcrossing

import "testing"

func TestHedgehogName(t *testing.T) {
	if ok := Hedgehog.Name() == hedgehog; !ok {
		t.Fatal("Hedgehog.Name() != hedgehog")
	}
}
