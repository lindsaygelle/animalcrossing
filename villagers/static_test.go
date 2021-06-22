package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestStaticAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Static.Animal() == s; !ok {
		t.Fatalf("%s != %s", Static.Animal(), s)
	}
}

func TestStaticName(t *testing.T) {
	if ok := Static.Name() == static; !ok {
		t.Fatalf("%s != %s", Static.Name(), static)
	}
}
