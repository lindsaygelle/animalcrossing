package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEuniceAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Eunice.Animal() == s; !ok {
		t.Fatalf("%s != %s", Eunice.Animal(), s)
	}
}

func TestEuniceName(t *testing.T) {
	if ok := Eunice.Name() == eunice; !ok {
		t.Fatalf("%s != %s", Eunice.Name(), eunice)
	}
}
