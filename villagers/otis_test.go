package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOtisAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Otis.Animal() == s; !ok {
		t.Fatalf("%s != %s", Otis.Animal(), s)
	}
}

func TestOtisName(t *testing.T) {
	if ok := Otis.Name() == otis; !ok {
		t.Fatalf("%s != %s", Otis.Name(), otis)
	}
}
