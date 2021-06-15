package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJittersName(t *testing.T) {
	if ok := Jitters.Name() == jitters; !ok {
		t.Fatalf("%s != %s", Jitters.Name(), jitters)
	}
}

func TestJittersSpecies(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Jitters.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jitters.Animal(), s)
	}
}
