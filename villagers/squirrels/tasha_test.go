package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTashaAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Tasha.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tasha.Animal(), s)
	}
}

func TestTashaName(t *testing.T) {
	if ok := Tasha.Name() == tasha; !ok {
		t.Fatalf("%s != %s", Tasha.Name(), tasha)
	}
}
