package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCallyAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Cally.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cally.Animal(), s)
	}
}

func TestCallyName(t *testing.T) {
	if ok := Cally.Name() == cally; !ok {
		t.Fatalf("%s != %s", Cally.Name(), cally)
	}
}
