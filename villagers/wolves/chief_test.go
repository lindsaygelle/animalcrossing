package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChiefAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Chief.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chief.Animal(), s)
	}
}

func TestChiefName(t *testing.T) {
	if ok := Chief.Name() == chief; !ok {
		t.Fatalf("%s != %s", Chief.Name(), chief)
	}
}
