package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLuluHippoAnimal(t *testing.T) {
	var s string = animals.Hippo.Name()
	if ok := LuluHippo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lulu.Animal(), s)
	}
}

func TestLuluHippoName(t *testing.T) {
	if ok := LuluHippo.Name() == lulu; !ok {
		t.Fatalf("%s != %s", LuluHippo.Name(), lulu)
	}
}
