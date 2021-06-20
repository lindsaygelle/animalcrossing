package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCandiAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Candi.Animal() == s; !ok {
		t.Fatalf("%s != %s", Candi.Animal(), s)
	}
}

func TestCandiName(t *testing.T) {
	if ok := Candi.Name() == candi; !ok {
		t.Fatalf("%s != %s", Candi.Name(), candi)
	}
}
