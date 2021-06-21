package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMerengueAnimal(t *testing.T) {
	var s string = animals.Rhinoceros.Name()
	if ok := Merengue.Animal() == s; !ok {
		t.Fatalf("%s != %s", Merengue.Animal(), s)
	}
}

func TestMerengueName(t *testing.T) {
	if ok := Merengue.Name() == merengue; !ok {
		t.Fatalf("%s != %s", Merengue.Name(), merengue)
	}
}
