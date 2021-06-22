package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSylvanaAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Sylvana.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sylvana.Animal(), s)
	}
}

func TestSylvanaName(t *testing.T) {
	if ok := Sylvana.Name() == sylvana; !ok {
		t.Fatalf("%s != %s", Sylvana.Name(), sylvana)
	}
}
