package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestElmerAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Elmer.Animal() == s; !ok {
		t.Fatalf("%s != %s", Elmer.Animal(), s)
	}
}

func TestElmerName(t *testing.T) {
	if ok := Elmer.Name() == elmer; !ok {
		t.Fatalf("%s != %s", Elmer.Name(), elmer)
	}
}
