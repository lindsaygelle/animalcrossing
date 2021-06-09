package animalcrossing

import "testing"

func TestAnteaterName(t *testing.T) {
	if ok := Anteater.Name() == anteater; !ok {
		t.Fatal("Anteater.Name() != anteater")
	}
}
