package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRexAnimal(t *testing.T) {
	var s string = animals.Lion.Name()
	if ok := Rex.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rex.Animal(), s)
	}
}

func TestRexName(t *testing.T) {
	if ok := Rex.Name() == rex; !ok {
		t.Fatalf("%s != %s", Rex.Name(), rex)
	}
}
