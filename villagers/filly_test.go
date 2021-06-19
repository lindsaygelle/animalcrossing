package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFillyAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Filly.Animal() == s; !ok {
		t.Fatalf("%s != %s", Filly.Animal(), s)
	}
}

func TestFillyName(t *testing.T) {
	if ok := Filly.Name() == filly; !ok {
		t.Fatalf("%s != %s", Filly.Name(), filly)
	}
}
