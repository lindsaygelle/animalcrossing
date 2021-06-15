package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPurrlAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Purrl.Animal() == s; !ok {
		t.Fatalf("%s != %s", Purrl.Animal(), s)
	}
}

func TestPurrlName(t *testing.T) {
	if ok := Purrl.Name() == purrl; !ok {
		t.Fatalf("%s != %s", Purrl.Name(), purrl)
	}
}
