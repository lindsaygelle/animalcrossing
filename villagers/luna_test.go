package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLunaAnimal(t *testing.T) {
	var s string = animals.Tapir.Name()
	if ok := Luna.Animal() == s; !ok {
		t.Fatalf("%s != %s", Luna.Animal(), s)
	}
}

func TestLunaName(t *testing.T) {
	if ok := Luna.Name() == luna; !ok {
		t.Fatalf("%s != %s", Luna.Name(), luna)
	}
}
