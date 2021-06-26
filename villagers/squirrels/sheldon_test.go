package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSheldonAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Sheldon.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sheldon.Animal(), s)
	}
}

func TestSheldonName(t *testing.T) {
	if ok := Sheldon.Name() == sheldon; !ok {
		t.Fatalf("%s != %s", Sheldon.Name(), sheldon)
	}
}
