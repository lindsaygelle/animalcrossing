package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBeardoName(t *testing.T) {
	if ok := Beardo.Name() == beardo; !ok {
		t.Fatalf("%s != %s", Beardo.Name(), beardo)
	}
}

func TestBeardoSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Beardo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Beardo.Animal(), s)
	}
}
