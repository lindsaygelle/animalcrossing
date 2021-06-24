package eagles

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKeatonAnimal(t *testing.T) {
	var s string = animals.Eagle.Name()
	if ok := Keaton.Animal() == s; !ok {
		t.Fatalf("%s != %s", Keaton.Animal(), s)
	}
}

func TestKeatonName(t *testing.T) {
	if ok := Keaton.Name() == keaton; !ok {
		t.Fatalf("%s != %s", Keaton.Name(), keaton)
	}
}
