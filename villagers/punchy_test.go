package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPunchyAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Punchy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Punchy.Animal(), s)
	}
}

func TestPunchyName(t *testing.T) {
	if ok := Punchy.Name() == punchy; !ok {
		t.Fatalf("%s != %s", Punchy.Name(), punchy)
	}
}
