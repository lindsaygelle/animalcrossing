package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFrillardAnimal(t *testing.T) {
	var s string = animals.FrillNeckedLizard.Name()
	if ok := Frillard.Animal() == s; !ok {
		t.Fatalf("%s != %s", Frillard.Animal(), s)
	}
}

func TestFrillardName(t *testing.T) {
	if ok := Frillard.Name() == frillard; !ok {
		t.Fatalf("%s != %s", Frillard.Name(), frillard)
	}
}
