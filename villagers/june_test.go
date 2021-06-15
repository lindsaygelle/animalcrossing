package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJuneAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := June.Animal() == s; !ok {
		t.Fatalf("%s != %s", June.Animal(), s)
	}
}

func TestJuneName(t *testing.T) {
	if ok := June.Name() == june; !ok {
		t.Fatalf("%s != %s", June.Name(), june)
	}
}
