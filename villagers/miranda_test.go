package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMirandaAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Miranda.Animal() == s; !ok {
		t.Fatalf("%s != %s", Miranda.Animal(), s)
	}
}

func TestMirandaName(t *testing.T) {
	if ok := Miranda.Name() == miranda; !ok {
		t.Fatalf("%s != %s", Miranda.Name(), miranda)
	}
}
