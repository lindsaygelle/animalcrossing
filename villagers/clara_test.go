package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestClaraAnimal(t *testing.T) {
	var s string = animals.Hippo.Name()
	if ok := Clara.Animal() == s; !ok {
		t.Fatalf("%s != %s", Clara.Animal(), s)
	}
}

func TestClaraName(t *testing.T) {
	if ok := Clara.Name() == clara; !ok {
		t.Fatalf("%s != %s", Clara.Name(), clara)
	}
}
