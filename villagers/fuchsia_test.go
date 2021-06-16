package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFuchsiaAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Fuchsia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Fuchsia.Animal(), s)
	}
}

func TestFuchsiaName(t *testing.T) {
	if ok := Fuchsia.Name() == fuchsia; !ok {
		t.Fatalf("%s != %s", Fuchsia.Name(), fuchsia)
	}
}
