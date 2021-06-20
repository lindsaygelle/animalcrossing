package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBorisAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Boris.Animal() == s; !ok {
		t.Fatalf("%s != %s", Boris.Animal(), s)
	}
}

func TestBorisName(t *testing.T) {
	if ok := Boris.Name() == boris; !ok {
		t.Fatalf("%s != %s", Boris.Name(), boris)
	}
}
