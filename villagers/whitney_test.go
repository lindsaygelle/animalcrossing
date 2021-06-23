package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWhitneyAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Whitney.Animal() == s; !ok {
		t.Fatalf("%s != %s", Whitney.Animal(), s)
	}
}

func TestWhitneyName(t *testing.T) {
	if ok := Whitney.Name() == whitney; !ok {
		t.Fatalf("%s != %s", Whitney.Name(), whitney)
	}
}
