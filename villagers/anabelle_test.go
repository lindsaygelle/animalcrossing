package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAnabelleName(t *testing.T) {
	if ok := Anabelle.Name() == anabelle; !ok {
		t.Fatalf("%s != %s", Anabelle.Name(), anabelle)
	}
}

func TestAnabelleSpecies(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Anabelle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Anabelle.Animal(), s)
	}
}
