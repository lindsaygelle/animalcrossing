package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAntonioName(t *testing.T) {
	if ok := Antonio.Name() == antonio; !ok {
		t.Fatalf("%s != %s", Antonio.Name(), antonio)
	}
}

func TestAntonioSpecies(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Antonio.Animal() == s; !ok {
		t.Fatalf("%s != %s", Antonio.Animal(), s)
	}
}
