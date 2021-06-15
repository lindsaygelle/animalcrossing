package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAlfonsoAnimal(t *testing.T) {
	var s string = animals.Alligator.Name()
	if ok := Alfonso.Animal() == s; !ok {
		t.Fatalf("%s != %s", Alfonso.Animal(), s)
	}
}

func TestAlfonsoName(t *testing.T) {
	if ok := Alfonso.Name() == alfonso; !ok {
		t.Fatalf("%s != %s", Alfonso.Name(), alfonso)
	}
}
