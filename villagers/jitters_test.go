package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJittersAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Jitters.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jitters.Animal(), s)
	}
}

func TestJittersName(t *testing.T) {
	if ok := Jitters.Name() == jitters; !ok {
		t.Fatalf("%s != %s", Jitters.Name(), jitters)
	}
}
