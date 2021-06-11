package animalcrossing

import "testing"

func TestPenguinName(t *testing.T) {
	if ok := Penguin.Name() == penguin; !ok {
		t.Fatal("Penguin.Name() != penguin")
	}
}
