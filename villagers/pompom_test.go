package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPompomAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Pompom.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pompom.Animal(), s)
	}
}

func TestPompomName(t *testing.T) {
	if ok := Pompom.Name() == pompom; !ok {
		t.Fatalf("%s != %s", Pompom.Name(), pompom)
	}
}
