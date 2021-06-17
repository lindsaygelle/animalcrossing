package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJambetteAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Jambette.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jambette.Animal(), s)
	}
}

func TestJambetteName(t *testing.T) {
	if ok := Jambette.Name() == jambette; !ok {
		t.Fatalf("%s != %s", Jambette.Name(), jambette)
	}
}
