package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPonchoName(t *testing.T) {
	if ok := Poncho.Name() == poncho; !ok {
		t.Fatalf("%s != %s", Poncho.Name(), poncho)
	}
}

func TestPonchoSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Poncho.Animal() == s; !ok {
		t.Fatalf("%s != %s", Poncho.Animal(), s)
	}
}
